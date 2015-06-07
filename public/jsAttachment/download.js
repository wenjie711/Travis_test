/**
 * Created by hebihui on 2015/6/4.
 */

/*-----------       视频加载       ---------------------*/
function setAttachmentUrl(url){
    //var video = "http://www.jplayer.org/video/m4v/Big_Buck_Bunny_Trailer.m4v";

    var video=url;
    $("#jquery_jplayer_1").jPlayer({
        ready: function () {
            $(this).jPlayer("setMedia", {
                title: "Big Buck Bunny",
                m4v: video,
                //m4v: "http://www.jplayer.org/video/m4v/Big_Buck_Bunny_Trailer.m4v",
                //ogv: "http://www.jplayer.org/video/ogv/Big_Buck_Bunny_Trailer.ogv",
                //webmv: "http://www.jplayer.org/video/webm/Big_Buck_Bunny_Trailer.webm",
                poster: "http://www.jplayer.org/video/poster/Big_Buck_Bunny_Trailer_480x270.png"
            });
        },
        swfPath: "../../dist/jplayer",
        supplied: "webmv, ogv, m4v",
        size: {
            width: "640px",
            height: "360px",
            cssClass: "jp-video-360p"
        },
        useStateClassSkin: true,
        autoBlur: false,
        smoothPlayBar: true,
        keyEnabled: true,
        remainingDuration: true,
        toggleDuration: true
    });
}

/*-----------------       绑定附件ID   ----------------------*/
var attachmentParam = {
    //"video":"http://7xiti1.com1.z0.glb.clouddn.com/test-clip.mp4",
    "video":"attachment",
    "pdf":"456",
    "ppt":"789"
}
function setAttachmentParam(attachmentParam){
    var video = attachmentParam.video;
    var pdf = attachmentParam.pdf;
    var ppt = attachmentParam.ppt;
    if(video!="0"){
       // 激活按钮,绑定ID
        $("#video").show();
        $("#video").attr("name",video);
    }
    if(pdf!="0"){
        // 激活按钮,绑定ID
        $("#")
    }
    if(ppt!="0"){
        // 激活按钮,绑定ID
    }
}


/*------------------   获取附件URL    ---------------------*/

function getUrl(attachmentId){
    //ajax
    $.get("http://localhost:9000/"+attachmentId, function(json){
        url=json.url;
        setAttachmentUrl(json.url);
    });

}

/*-------------------   预览附件    -----------------------*/

function preView(type){
    //显示相应模态框
}

/*------------------  下载附件     ------------------------*/

function downLoad(attachmentId){

}

/*------------------ 关闭视频      -------------------------*/

function closeVideo(){
    $("#jquery_jplayer_1").jPlayer("pause");
    $("#myModal3").show();
}