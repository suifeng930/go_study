package process

import "fmt"

//负责对在线用户列表做操作

//因为userMgr 再服务器端有且仅有一个，因此将其定义未一个全局变量
var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsersMap map[int]*UserProcess
}

//完成userMgr 初始化工作
func init() {
	userMgr = &UserMgr{
		onlineUsersMap: make(map[int]*UserProcess, 1024),
	}
}

//完成对onlineUser 的添加操作

func (this *UserMgr) AddOnlineUser(up *UserProcess) {
	this.onlineUsersMap[up.UserId] = up
}

// 删除onlineUser指定用户
func (this *UserMgr) DelOnlineUser(userId int) {
	delete(this.onlineUsersMap, userId)
}

//通过主键查询用户
func (this *UserMgr) GetOnlineUsersByUserId(userId int) (up *UserProcess, err error) {
	up, ok := this.onlineUsersMap[userId]
	if !ok {
		fmt.Println("你要查询的用户不在线，")
		err = fmt.Errorf("用户 %d 不存在，", userId)
		return
	}
	return
}

// 查询onlineUser

func (this *UserMgr) GetAllOnlineUsers() map[int]*UserProcess {
	return this.onlineUsersMap
}
