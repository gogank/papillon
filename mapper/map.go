package mapper

var linkMap map[string]string

func init(){
	linkMap = new(map[string]string)
}

func Get(key string) string {
	if hash,ok := linkMap[key];ok {
		return hash
	}
	return nil
}

func Put(key string,hash string) {
	//TODO need put link
}
