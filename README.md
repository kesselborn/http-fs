# http-fs

unsecure http server to get/put/delete files in a directory

# Usage

    $ ./http-fs -addr=localhost:8000 -dir=/tmp
    
    serving current /tmp at localhost:8000
    
    - get file /tmp/foo/bar:
    
        curl -O localhost:8000/tmp/foo/bar
    
    - delete file /tmp/foo/bar:
    
        curl -XDELETE localhost:8000/tmp/foo/bar
    
    - upload file to /tmp/foo/bar:
    
        curl -T /tmp/file localhost:8000/tmp/foo/bar
    



