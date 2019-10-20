package repo

import (
	"encoding/json"
	"fmt"
	"podkeeper/pkg/pod"
)

type PodRepository interface {
	add(pod *pod.Pod) (*pod.Pod, error)
	get(id string) (pod *pod.Pod, err error)
}

var dataStore = make(map[string]interface{})

type InMemoryPodRepository struct {
}

func (im *InMemoryPodRepository) add(pod *pod.Pod) (*pod.Pod, error) {
	podData, _ := json.Marshal(pod)
	dataStore[pod.Id] = podData
	return pod, nil
}

func (im *InMemoryPodRepository) get(id string) (pod *pod.Pod, err error) {
	podData := dataStore[id]
	podString := fmt.Sprintf("%v", podData)
	_ = json.Unmarshal([]byte(podString), pod)
	return pod, nil
}
