
# 🦇 rn-proxy🦇

application that creates a proxy to sit inbetween Exalt and the game servers and provide code hooks for mods/cheats, plugins and anything else   
**\+** has an extensive library of code injection, hooking and quality-of-life Exalt tools
 
- 🎮 works on Windows 10 ***and any new Macs!**
- 👓 overlayed mod-menu in-game and an external window for extra controls and settings 
- 📈 overlay your FPS, ping, memory usage counters and graphs in the minimap section
- ⌨ custom settings for keybinds and mouse events
- 🔃 auto-updating after game versions *(although it's safer to just wait til this is updated or do it manually)*
- 🔌 can hook the game socket and manipulated packets instead of using a proxy server  *[in progress]* 
- ⚡ load and save all Exalt settings to files   
- 🤯 write custom modules in Go , C/C++ and **FORTRAN!!**
- 💬 general game and packet/network logging options for game sessions
- 👥 custom client tokens (aka HWID spoofing for now)  

### Normal use
1. Download the version for your PC from the [releases]() page
2. Unzip it and drag the folder somewhere
3. Edit the `config/config.yaml` file or just leave it for default settings
4. Run the application file - open this before Exalt in the future to start it automatically  

##### Settings:
5. Exalt's settings will be saved in a file instead of the registry unless you turn it off the setting in `config.yaml`  	
		This means the first time you run the game *your settings while using the mod will be brand new*
6. The file can be edited, saved or copy pasted to save and then any file loaded, from Exalt itself or an account managerr
7. The files are saved in different places for each OS. but can be changed in the config anyway:
  - **Windows**: `C:\Users\username\Documents\RealmOfTheMadGod\UserData\PlayerPrefs.json`
#### 
 
### Building from source

#todo


