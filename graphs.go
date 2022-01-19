package main

//
//import "fmt"
//
////DependencyDAG represents a dependency directed graph
//type DependencyDAG struct {
//	Project string
//	Ins     []*DependencyDAG
//	Outs    []*DependencyDAG
//	Visited bool
//}
//
//func (d *DependencyDAG) IsSource() bool {
//	return len(d.Ins) > 0 && len(d.Outs) == 0
//}
//
//func (d *DependencyDAG) IsSink() bool {
//	return len(d.Ins) == 0 && len(d.Outs) > 0
//}
//
//func (d *DependencyDAG) IsSingular() bool {
//	return len(d.Ins) == 0 && len(d.Outs) == 0
//}
//
//func PrintDependencyOrder(projectDependencies [][]string) {
//	if len(projectDependencies) == 0 {
//		return
//	}
//	projects := createDependencyDAG(projectDependencies)
//
//	sources := findProjectSourcesAndSinks(projects)
//
//	for k, v := range sources {
//		var depsForSource []string = PrintBFS(k, v)
//		fmt.Println(depsForSource)
//	}
//}
//
////auxiliary
//func PrintBFS(k string, v *DependencyDAG) []string {
//	queue := make([]*DependencyDAG, 0)
//	enqueue(queue, v)
//	v.Visited = true
//
//	for len(queue) > 0 {
//		r := deque(queue)
//		//visit(r)
//		for _, node := range r.Outs {
//
//		}
//	}
//	//In fact, if what you want is a basic and easy to use fifo queue, slice provides all you need.
//	//
//	//	queue := make([]int, 0)
//	//// Push to the queue
//	//queue = append(queue, 1)
//	//// Top (just get next element, don't remove it)
//	//x = queue[0]
//	//// Discard top element
//	//queue = queue[1:]
//	//// Is empty ?
//	//if len(queue) == 0 {
//	//	fmt.Println("Queue is empty !")
//
//}
//
//func enqueue(src []*DependencyDAG, v *DependencyDAG) []*DependencyDAG {
//	return append(src, v)
//}
//
//func deque(src []*DependencyDAG) *DependencyDAG {
//	r := src[0]
//	src = src[1:]
//	return r
//}
//
//func findProjectSourcesAndSinks(projects map[string]*DependencyDAG) map[string]*DependencyDAG {
//	output := make(map[string]*DependencyDAG)
//	for k, v := range projects {
//		if v.IsSource() || v.IsSingular() {
//			output[k] = v
//		}
//	}
//	return output
//}
//
//func createDependencyDAG(dependencyPairs [][]string) map[string]*DependencyDAG {
//	// (a,b)-> a Outs(b) + b Ins(a)
//	projects := make(map[string]*DependencyDAG)
//
//	for _, pair := range dependencyPairs {
//		if pair[0] == "" || pair[1] == "" {
//			continue
//		}
//		_, ok0 := projects[pair[0]]
//		_, ok1 := projects[pair[1]]
//		// 0->1
//		//0->(out)    ->1(in)
//		if !ok0 && !ok1 {
//			projects[pair[0]] = &DependencyDAG{
//				Project: pair[0],
//			}
//			projects[pair[1]] = &DependencyDAG{
//				Project: pair[1],
//			}
//			projects[pair[0]].Outs = []*DependencyDAG{projects[pair[1]]}
//			projects[pair[1]].Ins = []*DependencyDAG{projects[pair[0]]}
//		} else if !ok0 && ok1 {
//			projects[pair[0]] = &DependencyDAG{
//				Project: pair[0],
//			}
//			projects[pair[0]].Outs = []*DependencyDAG{projects[pair[1]]}
//			projects[pair[1]].Ins = append(projects[pair[1]].Ins, projects[pair[0]])
//		} else if ok0 && !ok1 {
//			projects[pair[1]] = &DependencyDAG{
//				Project: pair[1],
//			}
//			projects[pair[0]].Outs = append(projects[pair[0]].Outs, projects[pair[1]])
//			projects[pair[1]].Ins = []*DependencyDAG{projects[pair[0]]}
//		} else {
//			projects[pair[0]].Outs = append(projects[pair[0]].Outs, projects[pair[1]])
//			projects[pair[1]].Ins = append(projects[pair[1]].Ins, projects[pair[0]])
//		}
//
//	}
//
//	return projects
//}
