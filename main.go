package main

import (
	"flag"
	"rpc-server/config"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

// NewConfig에 들어가는 toml path 값은 유동적으로 작성
// 플래그 이름, 기본값, 플래그 용도에 대한 간략한 설명. -h 또는 --help 출력 시 표시됨
// -config 또는 --config를 사용해 값을 전달할 수 있다.
// 플래그 기능은 CLI 사용자에게 가이드를 제공한다.
var configFlag = flag.String("config", "./config.toml", "config path")

func main() {
	// flag 패키지: Go에서 명령줄 플래그를 처리하는 표준 라이브러리 패키지.
	// `-config="./config.toml"`이라는 명령줄 플래그가 전달되면, 이 값을 플래그 변수에 저장한다.
	flag.Parse() // 정의한플래그들을 기준으로 CLI 입력값을 읽어 변수에 저장하도록 처리. 반드시 플래그 사용 전에 호출해야 한다.
	// flag.Parse() 호줄 전에는 플래그 변수에 기본값이 들어가 있음. (configFlag의 경우 "./config.toml")

	//fmt.Println(*configFlag) // configFlag에 저장된 값 출력

	// config 생성, 저장된 플래그 값(*configFlag: configFlag 변수의 주소 참조)을 넣어줌
	// CLI에서 go run . -config=test 명령을 입력하면 test가 path 값으로 입력됨.
	// 즉, 현재 경로의 test 파일/폴더를 읽는다.
	//지정하지 않으면 기본값인 ./config.toml이 입력된다.
	config.NewConfig(*configFlag)
}
