Set-Variable -Name "BINARY" -Value "swage"
Set-Variable -Name "IMAGE" -Value "swage"
Set-Variable -Name "REPO" -Value "markruler"
Set-Variable -Name "REPO_PORT" -Value ""
Set-Variable -Name "VERSION" -Value $(Get-Content -Path ./VERSION)

Write-Host "Clean..."
$OLD_IMAGES = docker images -q swage
foreach ($OLD in $OLD_IMAGES) {
  docker rmi --force $OLD
}
Remove-Item $(${BINARY} + "." + ${VERSION} + "-linux-amd64") -Force
Remove-Item *.xlsx -Force

$ErrorActionPreference = "Stop"

Write-Host "Build..."
$Env:GOOS="linux"; $Env:GOARCH="amd64"; go build -o $(${BINARY} + "." + ${VERSION} + "-linux-amd64") main.go

Write-Host "Make a image...${BINARY}:${VERSION}"
docker build `
--tag ${BINARY}:${VERSION} `
--tag ${BINARY}:latest `
--build-arg VERSION=${VERSION} `
--file ./Dockerfile `
${PWD}

Write-Host "Tag..."
docker tag ${BINARY}:${VERSION} $REPO$REPO_PORT/${IMAGE}:${VERSION}
docker tag ${BINARY}:${VERSION} $REPO$REPO_PORT/${IMAGE}:latest

# Write-Host "Push..."
# docker push $REPO$REPO_PORT/${IMAGE}:${VERSION}
# docker push $REPO$REPO_PORT/${IMAGE}:latest

Write-Host "Run..."
docker run `
--rm `
--interactive `
--tty `
--volume ${PWD}/examples/testdata:/testdata `
${IMAGE}:${VERSION} gen /testdata/json/editor.swagger.json `
--output /testdata/docker-swage.xlsx `
--verbose
