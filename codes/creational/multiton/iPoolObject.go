package multiton

type iPoolObject interface {
	getID() string //这是任何 ID，可用于比较两个不同的池对象
}
