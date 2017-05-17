package tasks

import "testing"
import "fmt"

// to run the test, go to models/tasks and then run "go test"
// testing is pkg that allows us to record errors
func TestCRUD(t *testing.T) {
	// connect to another server in another network
	sess, err := mgo.Dial("localhost:27017")
	if err != nil {
		t.Fatalf("error dialing Mongo: %v", err)
	}
	defer sess.Close()
	store := &MongoStore{
		Session:        sess,
		DatabaseName:   "test",
		CollectionName: "tasks",
	}

	newtask := &NewTask{
		Title: "Learn MongoDB",
		Tags:  []string("mongo", "info-344"),
	}
	// insert a task
	task, err := store.Insert(newtask)
	if err != nil {
		t.Errorf("error inserting new task:%v", err)
	}
	fmt.Println(task.ID)
	// retrieve a task
	task2, err := store.Get(task.ID)
	if err != nil {
		t.Errorf("error fetching task: %v", err)
	}
	if task2.Title != task.Title {
		t.Errorf("task title didn't match, expected %s but got %s", task.Title, task2.Title)
	}

	//remove all test Data
	sess.DB(store.DatabaseName).C(store.CollectionName).RemoveAll(nil)
}
