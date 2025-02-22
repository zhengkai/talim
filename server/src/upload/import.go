package upload

import "project/zj"

func importUser(ab []byte) {
	zj.J(`user`, len(ab))

}

func importTweet(ab []byte) {
	zj.J(`tweet`, len(ab))
}
