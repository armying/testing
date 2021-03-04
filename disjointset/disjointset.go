package disjointset


// DisjointSet is the interface for the disjoint-set (or union-find) data
// structure.
// Do not change the definition of this interface.
type DisjointSet interface {
	// UnionSet(s, t) merges (unions) the sets containing s and t,
	// and returns the representative of the resulting merged set.
	UnionSet(int, int) int
	// FindSet(s) returns representative of the class that s belongs to.
	FindSet(int) int
}

type DisjointsetManager struct {
    Collection map[int]*Dset //this keeps track of tree leaves
}

type Dset struct {
    Value int
    Parent *Dset
}

func (s *DisjointsetManager) FindSet(i int) int {
    switch s.Collection[i] {
    case nil:
        q := new(Dset)
        q.Parent = q
        q.Value = i
        s.Collection[i] = q
        return s.Collection[i].Value
    case s.Collection[i].Parent:
        return i
    default:
        s.Collection[i].Parent = s.Collection[s.FindSet(s.Collection[i].Parent.Value)]
        return s.Collection[i].Parent.Value
    }
}

func (s *DisjointsetManager) UnionSet(i int, j int) int {
    iparent := s.Collection[s.FindSet(i)]
    jparent := s.Collection[s.FindSet(j)]
    iparent.Parent = jparent
    return iparent.Parent.Value
}

// NewDisjointSet creates a struct of a type that satisfies the DisjointSet interface.
func NewDisjointSet() DisjointSet {
	//panic("TODO: implement this!")
    s := new(DisjointsetManager)
    s.Collection = make(map[int]*Dset)
    return s
}

