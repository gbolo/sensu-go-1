// +build integration,!race

package ring

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/sensu/sensu-go/backend/etcd"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	e, cleanup := etcd.NewTestEtcd(t)
	defer cleanup()

	client, err := e.NewClient()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	ring := EtcdGetter{Client: client, ClientID: "TestAdd"}.GetRing("testadd")
	if err := ring.Add(context.Background(), "foo"); err != nil {
		t.Fatal(err)
	}
}

func TestRemove(t *testing.T) {
	t.Parallel()

	e, cleanup := etcd.NewTestEtcd(t)
	defer cleanup()

	client, err := e.NewClient()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	ring := EtcdGetter{Client: client, ClientID: "TestRemove"}.GetRing("testremove")
	if err := ring.Add(context.Background(), "foo"); err != nil {
		t.Fatal(err)
	}
	if err := ring.Remove(context.Background(), "foo"); err != nil {
		t.Fatal(err)
	}
}

func TestNext(t *testing.T) {
	t.Parallel()

	e, cleanup := etcd.NewTestEtcd(t)
	defer cleanup()

	client, err := e.NewClient()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	ring := EtcdGetter{Client: client, ClientID: "TestNext"}.GetRing("testnext")

	items := []string{"foo", "bar", "baz"}
	for _, item := range items {
		if err := ring.Add(context.Background(), item); err != nil {
			t.Fatal(err)
		}
	}

	var got []string

	for i := 0; i < 9; i++ {
		item, err := ring.Next(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		got = append(got, item)
	}

	want := append(items, items...)
	want = append(want, items...)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("bad values: got %v, want %v", got, want)
	}

	if err := ring.Remove(context.Background(), "bar"); err != nil {
		t.Fatal(err)
	}

	newItems := []string{"foo", "baz"}
	want = want[:0]
	for i := 0; i < 5; i++ {
		want = append(want, newItems...)
	}

	got = got[:0]

	for i := 0; i < 10; i++ {
		item, err := ring.Next(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		got = append(got, item)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("bad values: got %v, want %v", got, want)
	}
}

func TestErrorOnNext(t *testing.T) {
	t.Parallel()

	e, cleanup := etcd.NewTestEtcd(t)
	defer cleanup()

	client, err := e.NewClient()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	getterA := EtcdGetter{Client: client, ClientID: "TestErrorOnNextA"}

	getterB := EtcdGetter{Client: client, ClientID: "TestErrorOnNextB"}

	r1 := getterA.GetRing("blocknext")
	r2 := getterB.GetRing("blocknext")

	if err := r1.Add(context.Background(), "foo"); err != nil {
		t.Fatal(err)
	}
	if err := r2.Add(context.Background(), "bar"); err != nil {
		t.Fatal(err)
	}

	value, err := r1.Next(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if got, want := value, "foo"; got != want {
		t.Fatalf("bad values: got %q, want %q", got, want)
	}
	_, err = r2.Next(context.Background())
	if err != ErrNotOwner {
		t.Fatalf("wanted ErrNotOwner, got %v", err)
	}

	value, err = r2.Next(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if got, want := value, "bar"; got != want {
		t.Fatalf("bad values: got %q, want %q", got, want)
	}

	_, err = r1.Next(context.Background())
	if err != ErrNotOwner {
		t.Fatalf("wanted ErrNotOwner, got %v", err)
	}
}

func TestTransferOwnership(t *testing.T) {
	t.Parallel()

	e, cleanup := etcd.NewTestEtcd(t)
	defer cleanup()

	client, err := e.NewClient()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	getterA := EtcdGetter{Client: client, ClientID: "TestTransferOwnerA"}
	getterB := EtcdGetter{Client: client, ClientID: "TestTransferOwnerB"}

	r1 := getterA.GetRing("testtransfer")
	r2 := getterB.GetRing("testtransfer")

	if err := r1.Add(context.Background(), "foo"); err != nil {
		t.Fatal(err)
	}
	if err := r2.Add(context.Background(), "foo"); err != nil {
		t.Fatal(err)
	}

	value, err := r2.Next(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if got, want := value, "foo"; got != want {
		t.Fatalf("bad values: got %q, want %q", got, want)
	}

	if _, err := r1.Next(context.Background()); err != ErrNotOwner {
		t.Fatalf("wanted ErrNotOwner, got %v", err)
	}
}

func TestErrNotOwner(t *testing.T) {
	t.Parallel()

	e, cleanup := etcd.NewTestEtcd(t)
	defer cleanup()

	client, err := e.NewClient()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	getter := EtcdGetter{Client: client, ClientID: "TestErrNotOwner"}

	r1 := getter.GetRing("testerrnotowner")
	r2 := getter.GetRing("testerrnotowner")
	r2.(*Ring).clientID = "something-else-entirely"

	if err := r1.Add(context.Background(), "foo"); err != nil {
		t.Fatal(err)
	}
	if got, want := r2.Remove(context.Background(), "foo"), ErrNotOwner; got != want {
		t.Fatalf("bad error: got %v, want %v", got, want)
	}
}

func TestExpire(t *testing.T) {
	t.Parallel()

	e, cleanup := etcd.NewTestEtcd(t)
	defer cleanup()

	client, err := e.NewClient()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	ring := EtcdGetter{Client: client, ClientID: "TestExpire"}.GetRing("testexpire").(*Ring)
	ring.leaseTimeout = 1

	if err := ring.Add(context.Background(), "foo"); err != nil {
		t.Fatal(err)
	}

	// Simulate the client dying
	if _, err := ring.client.Revoke(context.Background(), leaseIDCache[ring.Name]); err != nil {
		t.Fatal(err)
	}

	// Give the cluster some time to expire the lease. Unfortunately there
	// doesn't seem to be any way to be informed of when the lease expires.
	time.Sleep(time.Second * 5)

	_, err = ring.Next(context.Background())
	if got, want := err, ErrEmptyRing; got != want {
		t.Fatalf("bad error: got %v, want %v", got, want)
	}
}

func TestAddToEmptyRingAfterDelete(t *testing.T) {
	t.Parallel()

	e, cleanup := etcd.NewTestEtcd(t)
	defer cleanup()

	client, err := e.NewClient()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	ring := EtcdGetter{Client: client, ClientID: "TestAddToEmptyRingAfterDelete"}.GetRing("test_add_to_empty_ring_after_delete")

	items := []string{"foo", "bar"}
	for _, item := range items {
		if err := ring.Add(context.Background(), item); err != nil {
			t.Fatal(err)
		}
	}

	for _, item := range items {
		got, err := ring.Next(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		if want := item; got != want {
			t.Fatalf("bad values: got %q, want %q", got, want)
		}
	}

	for _, item := range items {
		if err := ring.Remove(context.Background(), item); err != nil {
			t.Fatal(err)
		}
	}

	_, err = ring.Next(context.Background())
	if got, want := err, ErrEmptyRing; got != want {
		t.Fatalf("bad error: got %v, want %v", got, want)
	}

	for _, item := range items {
		if err := ring.Add(context.Background(), item); err != nil {
			t.Fatal(err)
		}
	}

	for _, item := range items {
		got, err := ring.Next(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		if want := item; got != want {
			t.Fatalf("bad values: got %q, want %q", got, want)
		}
	}
}

func TestAddToRingWhenValueExists(t *testing.T) {
	t.Parallel()

	e, cleanup := etcd.NewTestEtcd(t)
	defer cleanup()

	client, err := e.NewClient()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	ring := EtcdGetter{Client: client, ClientID: "TestAddWhenValueExists"}.GetRing("test_add_when_value_exists")

	items := []string{"foo", "bar"}
	for _, item := range items {
		if err := ring.Add(context.Background(), item); err != nil {
			t.Fatal(err)
		}
	}

	for _, item := range items {
		got, err := ring.Next(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		if want := item; got != want {
			t.Fatalf("bad values: got %q, want %q", got, want)
		}
	}

	for _, item := range items {
		if err := ring.Add(context.Background(), item); err != nil {
			t.Fatal(err)
		}
	}

	for _, item := range items {
		got, err := ring.Next(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		if want := item; got != want {
			t.Fatalf("bad values: got %q, want %q", got, want)
		}
	}
}

func TestProblematicSequenceOfNexts(t *testing.T) {
	t.Parallel()

	e, cleanup := etcd.NewTestEtcd(t)
	defer cleanup()

	client, err := e.NewClient()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	ringA := EtcdGetter{Client: client, ClientID: "TestProblematicSequenceA"}.GetRing("test_prob_seq")
	ringB := EtcdGetter{Client: client, ClientID: "TestProblematicSequenceB"}.GetRing("test_prob_seq")

	if err := ringA.Add(context.Background(), "foo"); err != nil {
		t.Fatal(err)
	}
	if err := ringB.Add(context.Background(), "bar"); err != nil {
		t.Fatal(err)
	}
	if got, err := ringA.Next(context.Background()); err != nil {
		t.Fatal(err)
	} else if want := "foo"; got != want {
		t.Fatalf("bad value: got %q, want %q", got, want)
	}
	if _, err := ringB.Next(context.Background()); err != ErrNotOwner {
		t.Fatalf("expected ErrNotOwner, got %v", err)
	}
	if got, err := ringB.Next(context.Background()); err != nil {
		t.Fatal(err)
	} else if want := "bar"; got != want {
		t.Fatalf("bad value: got %q, want %q", got, want)
	}
	if _, err := ringA.Next(context.Background()); err != ErrNotOwner {
		t.Fatalf("bad error: got %v, want ErrNotOwner", err)
	}
	done := make(chan struct{})
	go func() {
		defer close(done)
		// ringB must now wait until ringA calls Next
		_, _ = ringB.Next(context.Background())
	}()
	// The call to ringA.Next should unblock ringB.Next
	if got, err := ringA.Next(context.Background()); err != nil {
		t.Fatal(err)
	} else if want := "foo"; got != want {
		t.Fatalf("bad value: got %q, want %q", got, want)
	}
	select {
	case <-done:
	case <-time.After(5 * time.Second):
		t.Fatal("test timed out")
	}
}
