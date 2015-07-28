FIRST
----------------
## /etc/hosts must be writable. 

I recommend making it owned and writable by group staff:

    sudo chown :staff /etc/hosts && sudo chmod g+w /etc/hosts


USAGE
----------------
<br />
<br />
![hosts workflow](https://raw.githubusercontent.com/antscript/AlfredWorkflow/master/res/hostsTool-2.png)
#### add domain to hosts (add '127.0.0.1 example.com' to hosts file)
<br />
<br />
![hosts workflow](https://raw.githubusercontent.com/antscript/AlfredWorkflow/master/res/hostsTool-3.png)
#### remove domain to hosts (change '127.0.0.1 example.com' to '# 127.0.0.1 example.com')
<br />
<br />
![hosts workflow](https://raw.githubusercontent.com/antscript/AlfredWorkflow/master/res/hostsTool-4.png)
#### add custom host and domain ( add '192.168.1.1 ex.com' to hosts file )
