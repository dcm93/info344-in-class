package tasks

//Store defines an abstract interface for a Task object store
//Interfaces: contract on set methods that the object will support if it implements the interface
// GO uses Duck Typing, if you implement all methods, you will be treated as of a given interface

// This interface will be shared by stores for mongo and postgres. So the idea is that the data is //turned into these stores and then the handlers send them in to the corresponding DB. therefore, //handlers should not care if the db is relational or mongo, for them the stores are equal before //they implement the same interface.

type Store interface {
	//Insert inserts a NewTask and
	//returns the fully-populated Task or an error
	Insert(newtask *NewTask) (*Task, error)
	Get(ID interface{}) (*Task, error)
	GetAll() (*[]Task, error)
	Update(task *Task) error
}
