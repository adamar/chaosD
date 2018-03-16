










## Proposed test list

 - high disk IO
 - use up inodes
 - use up disk space
 - use all file descriptors
 - increase network latency
 - block certain domains 
     by adding to hosts file
 - use large amounts of cpu
 - use large amounts of memory
 - set a process oom score 
     really high and then 
     use lots of memor 
 - use cgroups to constrain
     certain processes
     https://github.com/containerd/cgroups
 - use all ephemeral sockets?
 - start lots of processes and
     fill the process table
 - create lots of zombie processes
 - set incorrect system time
 

