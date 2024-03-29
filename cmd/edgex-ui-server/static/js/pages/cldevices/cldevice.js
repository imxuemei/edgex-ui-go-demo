/*******************************************************************************
 * Copyright © 2017-2018 VMware, Inc. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *
 * @author: Huaqiao Zhang, <huaqiaoz@vmware.com>
 *******************************************************************************/
$(document).ready(function(){
	$("#get_cldevice_name").on('click',function(event){
		event.stopPropagation();
		var cldname = getCldeviceName()
		$("#cldevice_name").text(cldname)
	});
	$("#start_simpledevice").on('click',function(event){
		event.stopPropagation();
		startSimpledevice()
	});
	$("#stop_simpledevice").on('click',function(event){
		event.stopPropagation();
		stopSimpledevice()
	});
});

function getCldeviceName() {
	var name;
	$.ajax({
		url:'/api/v1/cldevices/testgetname',
		type:'GET',
		async: false,
		dataType: "text",
		success:function(data){
			name = data
		}
	});
	return name;
};

function startSimpledevice() {
	$.ajax({
		url:'/api/v1/cldevices/startsimpledevice',
		type:'GET',
		async: false,
		dataType: "text",
		success:function(data){
			$("#simpledevice_status").text("状态：已启动")
		}
	});
};

function stopSimpledevice() {
	$.ajax({
		url:'/api/v1/cldevices/stopsimpledevice',
		type:'GET',
		async: false,
		dataType: "text",
		success:function(data){
			$("#simpledevice_status").text("状态：已关闭")
		}
	});
}
