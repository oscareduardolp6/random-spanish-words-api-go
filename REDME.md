# Random Spanish Words API - Go 
Simple API that return a `json` array of *random* **spanish** words.
The default returns only one word in the array, but you can specify the number of words returned . 

# Usage 
## Dev 
### Dev Container 
If you are using *Visual Studio Code* and have *Docker*, the project comes with all the **Dev Container** files, you only need to run the [Dev Container Extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack) 

> For more information about the dockerfile image you can check [this post](https://medium.com/@quentin.mcgaw/ultimate-go-dev-container-for-visual-studio-code-448f5e031911)

## "Production" 
If you want to use it in "production" you can *compile* the project with **Go(1.13^)** using the commands : 
`go build` 
that will create a executable containing the server. 