//Program to fetch urls recursively and concurrently
package main

type Fetcher interface {
	//It return body of a url and urls found on a given url
	Fetch(url string) (body string, urls []string, err error)
}

//Crawl uses Fetcher to recursively crawl
// urls to a given depth
func Crawl(url string, depth int, fetcher Fetcher) {
	//fetch urls in parallel
	//don't fetch same url twice
	if depth <= 0 {
		fmt.Printf("Reached leaf url %s \n", url)
		return 
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	//Crawl child pages
	for _, u := urls {
		Crawl(u, depth - 1, tetcher)
	}
}

func main(){
	Crawl("http://golang.org/",4, interface{})
}

