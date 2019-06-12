package robotname

import (
	"errors"
	"fmt"
	"math/rand"
)

var used = map[string]bool{"": true}

//Robot only have names. They aren't very useful machines!
type Robot struct {
	name string
}

//Name gives returns a robot's name if it has one, otherwise returns a random name.
func (robot *Robot) Name() (name string, err error) {
	if robot.name != "" {
		return robot.name, nil
	}
	if len(used) >= 676000 {
		return "", errors.New("Exhausted namespace")
	}
	return robot.Reset(), nil
}

//Reset resets a robots name to a random configuration.
func (robot *Robot) Reset() string {
	robot.name = randomName(robot)
	return robot.name
}

func randomName(robot *Robot) string {
	name := ""
	for used[name] {
		letters := string(rand.Intn(26)+'A') + string(rand.Intn(26)+'A')
		numbers := rand.Intn(1000)
		name = fmt.Sprintf("%s%03d", letters, numbers)
	}

	used[name] = true
	robot.name = name
	return name
}
