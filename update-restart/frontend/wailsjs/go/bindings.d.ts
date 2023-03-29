export interface go {
  "main": {
    "App": {
		Greet(arg1:string):Promise<string>
    },
    "Updater": {
		CheckForUpdate():Promise<string>
    },
  }

}

declare global {
	interface Window {
		go: go;
	}
}
