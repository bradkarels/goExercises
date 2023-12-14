package graphs

import (
	"fmt"
	"testing"
)

func TestGraph_addDirectedEdge(t *testing.T) {
	//var tstNodez = make(map[string]map[string]void)
	graph := Graph{nodez: make(map[string]map[string]void)}
	graph.addDirectedEdge(`a`, `b`)
	graph.addDirectedEdge(`a`, `c`)
	graph.addDirectedEdge(`a`, `b`)
	graph.addDirectedEdge(`a`, `z`)
	//fmt.Println(graph.nodez)

	type fields struct {
		nodez map[string]map[string]void
	}
	type args struct {
		src string
		dst string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: `b/c why not...`,
			fields: fields{
				//nodez: tstNodez,
				//nodez: make(map[string]map[string]void),
				nodez: graph.nodez,
			},
			args: args{
				src: `a`,
				dst: `b`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				nodez: tt.fields.nodez,
			}
			g.addDirectedEdge(tt.args.src, tt.args.dst)
			if edge, exists := g.nodez[tt.args.src]; !exists {
				fmt.Println(edge)
				t.Errorf("Added Edge did not get added...")
			}
		})
	}
}

func TestGraph_removeDirectedEdge(t *testing.T) {
	graph := Graph{nodez: make(map[string]map[string]void)}
	graph.addDirectedEdge(`a`, `b`)
	graph.addDirectedEdge(`a`, `c`)
	graph.addDirectedEdge(`a`, `d`)
	_, b := graph.nodez[`aaa`]
	fmt.Printf("x: %t\n", b)
	fmt.Println(graph.nodez)
	type fields struct {
		nodez map[string]map[string]void
	}
	type args struct {
		src string
		dst string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: `delete tester`,
			fields: fields{
				nodez: graph.nodez,
			},
			args: args{
				src: `a`,
				dst: `d`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				nodez: tt.fields.nodez,
			}
			g.removeDirectedEdge(tt.args.src, tt.args.dst)
			fmt.Println(g.nodez[tt.args.src])
		})
	}
}
