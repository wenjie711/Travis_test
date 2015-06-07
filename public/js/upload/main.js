$(function() {
    var uploader = Qiniu.uploader({
        runtimes: 'html5,flash,html4',
        browse_button: 'pickfiles', //上传按钮button的ID
        container: 'container', //上传区域的DOM ID
        drop_element: 'container', //拖拽上传区域的ID
        max_file_size: '100mb', //最大允许上传文件的大小
        flash_swf_url: 'public/js/upload/Moxie.swf',
        dragdrop: true, //是否允许拖拽上传
        chunk_size: '4mb', //分片上传的大小
        uptoken_url: '/sharing/upToken', //请求的uptoken url,会在初始化的时候由qiniu sdk对该url进行一次请求获取token
        domain: '7xjf2x.com1.z0.glb.clouddn.com', //我们的七牛云存储空间
        unique_names: true, //是否由js端sdk自动生成key,这里开启
        auto_start: true, //选择文件后自动开始
        init: {
            'FilesAdded': function(up, files) {
                $('table').show();
                $('#success').hide();
                plupload.each(files, function(file) {
                    var progress = new FileProgress(file, 'fsUploadProgress');
                    progress.setStatus("等待...");
                });
            },
            'BeforeUpload': function(up, file) {
                var progress = new FileProgress(file, 'fsUploadProgress');
                var chunk_size = plupload.parseSize(this.getOption('chunk_size'));
                if (up.runtime === 'html5' && chunk_size) {
                    progress.setChunkProgess(chunk_size);
                }
                $("#submitBtn").attr("disabled","disabled");
            },
            'UploadProgress': function(up, file) {
                var progress = new FileProgress(file, 'fsUploadProgress');
                var chunk_size = plupload.parseSize(this.getOption('chunk_size'));
                progress.setProgress(file.percent + "%", up.total.bytesPerSec, chunk_size);

            },
            'UploadComplete': function() {
                $('#success').show(5000).hide();
                $("#submitBtn").removeAttr("disabled");
            },
            'FileUploaded': function(up, file, info) {
                $.ajax({
                    type:"POST",
                    contentType: "application/json",
                    url: "/sharing/upload",
                    data:info,
                    success: function() {
                        var progress = new FileProgress(file, 'fsUploadProgress');
                        progress.setComplete(up, info);
                    },
                });
            },
            'Error': function(up, err, errTip) {
                    $('table').show();
                    var progress = new FileProgress(err.file, 'fsUploadProgress');
                    progress.setError();
                    progress.setStatus(errTip);
                }
                // ,
                // 'Key': function(up, file) {
                //     var key = "";
                //     // do something with key
                //     return key
                // }
        }
    });

    uploader.bind('FileUploaded', function() {
        console.log('hello man,a file is uploaded');
    });
    $('#container').on(
        'dragenter',
        function(e) {
            e.preventDefault();
            $('#container').addClass('draging');
            e.stopPropagation();
        }
    ).on('drop', function(e) {
        e.preventDefault();
        $('#container').removeClass('draging');
        e.stopPropagation();
    }).on('dragleave', function(e) {
        e.preventDefault();
        $('#container').removeClass('draging');
        e.stopPropagation();
    }).on('dragover', function(e) {
        e.preventDefault();
        $('#container').addClass('draging');
        e.stopPropagation();
    });

});