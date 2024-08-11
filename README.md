# ffmpeg-video-modules

<br>

### 핵심 문제 해결 전략 🧐

* 아티클 1
* 아티클 2

<br><br>

### FFmpeg 및 FFprobe 설치 가이드 (Windows 10)

```
# 📌 FFmpeg 및 FFprobe 설치
  - [FFmpeg 다운로드 페이지](https://ffmpeg.org/download.html)로 이동
  - Windows 빌드 선택 (권장: [gyan.dev](https://www.gyan.dev/ffmpeg/builds/))
  - "Release builds"에서 `ffmpeg-release-essentials.zip` 다운로드
  - ZIP 파일을 `C:\ffmpeg`에 압축 해제


# 📌 PATH 환경 변수 추가
  - `C:\ffmpeg\bin` 경로 복사
  - 내 PC → 속성 → 고급 시스템 설정 → 환경 변수 클릭
  - 시스템 변수에서 `Path` 편집 → 새 경로에 `C:\ffmpeg\bin` 추가
  - 모든 창 닫기


# 📌 설치 확인
  - Win + R → "cmd" 입력 후 실행
  - `ffmpeg -version` 및 `ffprobe -version` 입력해 설치 확인
```

* 위 과정을 진행해야 Local 환경에서 ffmpeg 라이브러리를 사용할 수 있습니다.

<br><br>

### MariaDB 컨테이너 실행

```
docker network create --driver bridge mynetwork

docker network ls

docker run -d --name mariadb -p 3310:3306 -v mysql_db:/var/lib/mysql --network mynetwork -e MYSQL_DATABASE=ffmpeg-video-database -e MYSQL_ROOT_PASSWORD=password mariadb:latest
```

<br><br>

### 실행

```
# 📌 GoLand IDE 기준
Program arguments : -c deploy/dev/config.toml


# 📌Terminal 기준

```

<br><br>

### 테스트 코드 실행 시키기

```
# 📌 특정 디렉토리의 모든 테스트 코드 실행 [Git Bash 기준]
go test ./cmd/...


# 📌 커버리지 프로파일 생성 후, HTML 보고서 생성 [Git Bash 기준]
go test -v -coverprofile=coverage.out ./cmd/...
go tool cover -html=coverage.out
```

<br><br>

### 헬스 체크 API 테스트

```git bash
$ curl --http1.1 http://localhost:3031/api/health
```

<br><br>

### 테스트 비디오 생성 ffmpeg 명령어

```
ffmpeg -f lavfi -i color=c=blue:s=320x240:d=5 -vf "drawtext=fontfile='C\\Windows\\Fonts\\arial.ttf': text='Test Video 1': fontcolor=white: fontsize=24: x=(w-text_w)/2: y=(h-text_h)/2" -c:v libx264 -t 5 -pix_fmt yuv420p "C:\test_video_1.mp4"

ffmpeg -f lavfi -i color=c=red:s=320x240:d=6 -vf "drawtext=fontfile='C\\Windows\\Fonts\\arial.ttf': text='Test Video 2': fontcolor=white: fontsize=24: x=(w-text_w)/2: y=(h-text_h)/2" -c:v libx264 -t 6 -pix_fmt yuv420p "C:\test_video_2.mp4"

ffmpeg -f lavfi -i color=c=green:s=320x240:d=7 -vf "drawtext=fontfile='C\\Windows\\Fonts\\arial.ttf': text='Test Video 3': fontcolor=white: fontsize=24: x=(w-text_w)/2: y=(h-text_h)/2" -c:v libx264 -t 7 -pix_fmt yuv420p "C:\test_video_3.mp4"

ffmpeg -f lavfi -i color=c=yellow:s=320x240:d=5 -vf "drawtext=fontfile='C\\Windows\\Fonts\\arial.ttf': text='Test Video 4': fontcolor=black: fontsize=24: x=(w-text_w)/2: y=(h-text_h)/2" -c:v libx264 -t 5 -pix_fmt yuv420p "C:\test_video_4.mp4"

ffmpeg -f lavfi -i color=c=purple:s=320x240:d=6 -vf "drawtext=fontfile='C\\Windows\\Fonts\\arial.ttf': text='Test Video 5': fontcolor=white: fontsize=24: x=(w-text_w)/2: y=(h-text_h)/2" -c:v libx264 -t 6 -pix_fmt yuv420p "C:\test_video_5.mp4"
```
* 위 명령어를 실행하면, 테스트 비디오 파일들은 `C 드라이브`의 루트에 생성됩니다.

<br><br>

### Application Server Architecture

<img src="docs/server-architecture.png" alt="Application Server Architecture" width="800"/>

📌 [참고 Link](https://github.com/bxcodec/go-clean-arch)

<br><br>

### Go Clean Architecture 기반 폴더 구조

<img src="docs/folder-structure.png" alt="Folder Structure" width="250"/>

<br><br>

### ERD(Entity Relationship Diagram)

<img src="docs/erd.png" alt="ERD Diagram" width="600"/>

<br><br>

### Git 커밋 메시지 규칙

| Tag        | Description                                         |
|------------|-----------------------------------------------------|
| `feat`     | 새로운 기능을 추가한 경우 사용합니다.                               |
| `fix`      | 버그를 수정한 경우 사용합니다.                                   |
| `refactor` | 코드 리팩토링한 경우 사용합니다.                                  |
| `style`    | 코드 형식, 정렬, 주석 등의 변경(동작에 영향을 주는 코드 변경 없음)한 경우 사용합니다. |
| `test`     | 테스트 추가, 테스트 리팩토링(제품 코드 수정 없음, 테스트 코드에 관련된 모든 변경에 해당)한 경우 사용합니다.                                             |
| `docs`     | 문서를 수정(제품 코드 수정 없음)한 경우 사용합니다.                                             |
| `chore`    | 빌드 업무 수정, 패키지 매니저 설정 등 위에 해당되지 않는 모든 변경(제품 코드 수정 없음)일 경우 사용합니다.                                             |

<br><br>
