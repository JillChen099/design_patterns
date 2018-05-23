package observer

import "testing"

func TestExampleObserver(t *testing.T) {
	acc := NewConcreteAllyControlCenter("金庸群侠传")
	player1 := Player{"杨过"}
	acc.join(&player1)
	player2 := Player{"令狐冲"}
	acc.join(&player2)
	player3 := Player{"张无忌"}
	acc.join(&player3)
	player4 := Player{"段誉"}
	acc.join(&player4)
	player1.beAttacked(acc)

}