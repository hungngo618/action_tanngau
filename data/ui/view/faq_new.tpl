{{define "content"}}
<div class="jumbotron mb-0 pb-5">
    <div class="container">
        <div class="col-sm-12 col-center-block mb-5">
            <h1 align="center" class="text-primary">New FAQ</h1>
        </div>

        <div class="row justify-content-md-center" style="min-height: 400px;">
            <form class="form col-sm-10">

                <div class="form-group row">
                    <label class="col-2 col-form-label text-center">Intent Name</label>
                    <div class="col-10">
                        <input id="intent" class="form-control" type="text" placeholder="Enter intent name">
                    </div>
                </div>

                <div class="form-group row">
                    <label class="col-2 col-form-label text-center">Trinh Answer</label>
                    <div class="col-10">
                        <textarea name="answer" class="form-control" rows="5" id="answer1" placeholder="Enter answer"></textarea>
                    </div>
                </div>

                <div class="form-group row">
                    <label class="col-2 col-form-label text-center">Bom Answer</label>
                    <div class="col-10">
                        <textarea name="answer" class="form-control" rows="5" id="answer2" placeholder="Enter answer"></textarea>
                    </div>
                </div>

                <div class="form-group row">
                    <label class="col-2 col-form-label text-center">Samples</label>
                    <div class="row col-10 mr-0 pr-0" id="samples">
                        <div class="col-12 row  mx-0 px-0 my-1">
                            <div class="col-11 mr-0 pr-1 py-0">
                                <input type="text" class="form-control col-sm-12 sample" placeholder="Enter question">
                            </div>
                            <div class="col-1 mx-0 my-0 px-0 py-0">
                                <button type="button" data-id="1" class="form-control btn btn-md btn-danger fa fa-times deletesample">
                                </button>
                            </div>
                        </div>
                    </div>
                </div>

                <div class="form-group row">
                    <div class="row col-md-5 offset-md-7 mr-0 pr-0">
                        <div class="col-md-6 push-md-6">
                            <button type="button" class="fa fa-plus-square col-sm-12 btn btn-lg btn-success" id="add"> Add</button>
                        </div>

                        <div class="col-md-6 pull-md-6">
                        </div>

                    </div>
                </div>

                <div class="form-group row">
                    <div class="row col-md-5 offset-md-7 mr-0 pr-0">

                        <div class="col-md-6 push-md-6">
                            <button type="button" class="col-sm-12 btn btn-lg btn-primary" id="save">
                            <i id="loading" class="fa fa-floppy-o mr-2" style="font-size:24px; color:white;"></i>
                            Save</button>
                        </div>

                        <div class="col-md-6 pull-md-6">
                            <button type="button" class="col-sm-12 btn btn-default btn-lg" id="cancel"> 
                            <i id="loading2" class="fa fa-home mr-2" style="font-size:24px; color:black;"></i>
                            Cancel</button>
                        </div>

                    </div>
                </div>

            </form>
        </div>
    </div>
</div>
<script>
    $(document).ready(function() {
        $("#save").click(function() {


            var data = $("#samples input");
            var intent = $("#intent").val();
            // var instance = $("input:radio[name ='optradio']:checked").val();
            var trinhAnswer = $("#answer1").val();
            var bomAnswer = $("#answer2").val();

            var samples = [];
            $.each(data, function(i, v) {
                var val = JSON.stringify(v.value, escape);
                if (v.value != "") {
                    samples.push('{\"id\": 0, \"question\":' + val + ', \"intent_id\":0}');
                }
            });
            // trinhAnswer = JSON.stringify(trinhAnswer, escape);
            // bomAnswer = JSON.stringify(bomAnswer, escape);

            var dataStr = '{' +
                '"id": 0, ' +
                '"intent":"' + intent + '",' +
                '"trinh_answer":"' + trinhAnswer + '",' +
                '"bom_answer":"' + bomAnswer + '",' +
                '"questions":[' + samples + ']' +
                '}';
            if (samples.length == 0 || intent == '' || trinhAnswer == '' || bomAnswer == '') {
                alert("Please Fill All Fields");
            } else {
                $("#save").attr("disabled", true);
                $("#loading").removeClass("fa fa-floppy-o");
                $("#loading").addClass("fa fa-spinner fa-spin");
                var status = "";
                $.ajax({
                    type: "POST",
                    url: "/newfaq",
                    data: dataStr,
                    cache: false,
                    success: function(result) {
                        status = $.trim(result);
                        if (status == "success") {
                            location.href = "/faq";
                        } else {
                            alert(result);
                        }

                        $("#save").attr("disabled", false);
                        $("#loading").removeClass("fa fa-spinner fa-spin");
                        $("#loading").addClass("fa fa-floppy-o");
                    }
                });

            }
            return false;
        });

        $('body').on('click', '.deletesample', function() {
            var fID = $(this).data("id");
            var dataString = 'id=' + fID;
            element = $(this).parents(".row").get(0);
            //alert(element);

            $(element).remove();
        });

        $("#cancel").click(function() {
            location.href = "/faq";
            return false;
        });

        var cnt = 1;
        $(document).on('keyup', '.sample', function(e) {
            var key = e.which;
            if (key == 13) {
                cnt++;
                $("#samples").append(
                    '<div class="col-12 row  mx-0 px-0 my-1">' +
                    '<div class="col-11 mr-0 pr-1 py-0">' +
                    '<input type="text" class="form-control col-sm-12 sample" id="sample' + cnt + '" placeholder="Enter question">' +
                    '</div>' +
                    '<div class="col-1 mx-0 my-0 px-0 py-0">' +
                    '<button type="button" data-id="' + cnt + '" class="form-control btn btn-danger fa fa-times deletesample"></button>' +
                    '</div></div>');

                var idfocus = "#sample" + cnt;
                $(idfocus).focus();
            }
            return false;
        });

        $("#add").click(function() {
            cnt++;
            $("#samples").append(
                '<div class="col-12 row  mx-0 px-0 my-1">' +
                '<div class="col-11 mr-0 pr-1 py-0">' +
                '<input type="text" class="form-control col-sm-12 sample" id="sample' + cnt + '" placeholder="Enter question">' +
                '</div>' +
                '<div class="col-1 mx-0 my-0 px-0 py-0">' +
                '<button type="button" data-id="' + cnt + '" class="form-control btn btn-danger fa fa-times deletesample"></button>' +
                '</div></div>');
            var idfocus = "#sample" + cnt;
            $(idfocus).focus();
            return false;
        });

    });
</script>

{{end}}