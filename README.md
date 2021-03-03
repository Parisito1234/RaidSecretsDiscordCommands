# RaidSecretsDiscordCommands
YAGPDB.xyz custom commands for the RaidSecrets Discord server

## File Naming
Folder: Custom Command = `cc_{Command Group Name}` 

File: Custom Command = `cc_{Command Trigger}` 

* Copy the contents of the file into the text field for the custom command.
* Each file with the above previx is a command.


Folder: `setdb_{Associated Command/Group}`

File: `setdb_{Associated Command/Group}` 

*Follow these steps to use setdb commands:*
1. Create a _mod-only_ command or command group with this command inside it.
2. Set the command name to something arbitrary - I use `setdb`. This will be deleted when done.
3. Run the command once and it will return the associated message with formatted content.
	1. You can edit the command at this step for formatting or content purposes.
	2. Make sure you always run the command after saving it to properly update the database.
4. If the returned message displays correctly, replace the contents of the command with the next `setdb_{...}` file. Repeat steps 2 and 3.
5. After all `setdb_{...}` commands have been run and verified, delete your temp command.
6. Test the associated command with all args to check all variants and database entries.
**If a command in this repository gets updated, it will be done so in a way that requires a full overwrite of the previous command. All above steps will be followed to update.**


## To Do
- [ ] Add config, dependency, and override notes to each cc/group
- [ ] Create a sticky note command for `#datamining-discussion`
- [ ] Create system for multi-page command responses
- [ ] Create `-notes` command for moderation purposes outside of `-warn`

*Blah blah blah i'll fill the rest of this out later*
