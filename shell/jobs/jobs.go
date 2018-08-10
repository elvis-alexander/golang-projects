package jobs

import(
	"fmt"
	"bytes"
	"os/exec"
)

/*upper-case public, lower-case private*/
type job struct {
	next     *job
	prev     *job
	pid      int
	command  string
	exitcode int
	CmdObj *exec.Cmd
}


type List struct {
	head    *job
	tail    *job
	numJobs int
}

func NewList() *List {
	return &List{}
}

func (this *List) Add(pid int, command string, cmdObj *exec.Cmd) {
	var j = &job{pid: pid, command: command, CmdObj: cmdObj}
	if this.head == nil {
		/*empty list*/
		this.head = j
		this.tail = j
	} else {
		this.head.prev = j
		j.next = this.head
		this.head = j
	}
	this.numJobs += 1
}

/*returns */
func (this *List) Get(pid int) *job {
	var temp = this.head
	for temp != nil {
		if temp.pid == pid {
			return temp
		}
		temp = temp.next
	}
	return nil
}

func (this *List) Remove(pid int) bool {
	var ref = this.Get(pid)
	if ref == nil {
		return false
	}
	this.numJobs -= 1
	return true
}

func (this *List) Print() string {
	var b bytes.Buffer
	b.WriteString("PID\tCMD\n")
	var temp = this.head
	for temp != nil {
		b.WriteString(fmt.Sprintf("%d\t%s\n", temp.pid, temp.command))
		temp = temp.next
	}
	return b.String()
}

func (this *List) RemoveAll() {
	var temp = this.head
	for temp != nil {
		var next = temp.next
		/**temp = nil*/
		temp = next
	}
	this.head = nil
	this.numJobs = 0
}