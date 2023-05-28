package graph

import "fmt"

type Vertex struct {
	ID       int
	adjacent []*Vertex
}

type Graph struct {
	vertices []*Vertex
}

func New() *Graph {
	return &Graph{[]*Vertex{}}
}

func (g *Graph) AddVertex(ID int) error {
	if vertex := BSearch(g.vertices, ID); vertex != nil {
		return fmt.Errorf("vertex %v already exist", ID)
	}

	g.vertices = append(g.vertices, &Vertex{ID: ID})

	return nil
}

func (g *Graph) AddEdge(to, from int) error {
	toVertex := BSearch(g.vertices, to)
	if toVertex == nil {
		return fmt.Errorf("to argument need to exist in the graph")
	}

	fromVertex := BSearch(g.vertices, from)
	if fromVertex == nil {
		return fmt.Errorf("to argument need to exist in the graph")
	}

	if adjacent := BSearch(fromVertex.adjacent, to); adjacent != nil {
		return fmt.Errorf("edge from vertex %d ---> %d already exists", fromVertex.ID, toVertex.ID)
	}

	fromVertex.adjacent = append(fromVertex.adjacent, toVertex)

	return nil
}

func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("%d -> ", v.ID)

		adl := len(v.adjacent) - 1
		for adx, v := range v.adjacent {
			if adl == adx {
				fmt.Printf("%d ", v.ID)
			} else {
				fmt.Printf("%d -> ", v.ID)
			}
		}

		fmt.Println()
	}
}
