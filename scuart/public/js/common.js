$(function(){
    $('#powerPanel').on('change', function(){
        console.log('changed', $(this).val())
        send('/power/0', {on: true})
    })

    $('#btn-cct-group').on('click', function(){
        send('/cct/0', {cct: $('#cctGroup').val()})
    })
    $('#btn-dim-group').on('click', function(){
        send('/dim/0', {dim: $('#dimGroup').val()})
    })

    //operations for light1
    $('#powerLight1').on('change', function(){
        console.log('light1 power changed to', $(this).val())
        send('/power/1', {on: true})
    })
    $('#btn-dim-light1').on('click', function(){
        send('/dim/1', {dim: $('#dimLight1').val()})
    })
    $('#btn-cct-light1').on('click', function(){
        send('/cct/1', {cct: $('#cctLight1').val()})
    })
    //operations for light2   
    $('#powerLight2').on('change', function(){
        console.log('light2 power changed to', $(this).val())
        send('/power/2', {on: true})
    })
    $('#btn-dim-light2').on('click', function(){
        send('/dim/2', {dim: $('#dimLight2').val()})
    })
    $('#btn-cct-light2').on('click', function(){
        send('/cct/2', {cct: $('#cctLight2').val()})
    })

    $('#statusLight1').on('click', function(){
        updateLight(ls1)
    })
    $('#statusLight2').on('click', function(){
        updateLight(ls2)
    })

    getStatus()
})

var ls1 = null, ls2 = null // light status
function getStatus(){
    var ws = new WebSocket('ws://'+location.host+'/light/status')

    ws.onopen = function(e){
        console.log('connected')
        ws.send('connected')
    }
    ws.onclose = function(e){
        console.log('connetion closed')
    }
    ws.onerror = function(e){
        console.log('connetion error')
    }
    ws.onmessage = function(e){
        data = JSON.parse(e.data)
        console.log('data received', data)

        if(data.id == 1){
            ls1 = data
        }else  if(data.id == 2){
            ls2 = data
        }
    }
}

function updateLight(data){
    if (data == null) return

    $('#statusInputPower').val(data.inputPower)
    $('#statusDim').val(data.dim)
    $('#statusOuptputVoltage').val(data.outputVoltage)
    $('#statusOutputCurrent').val(data.outputCurrent)
    $('#statusTime').val(data.timestamp)
}

function send(url, data){
    console.log(url, data)
    $.ajax({
        type: 'POST',
        url: url,
        data: data,
        success: function(resp){
            swal('Done', getDescName(data)+' operation done successfully', 'success')
        }
    })
}

function getDescName(data){
    if(data.hasOwnProperty('on')) return 'power'
    if(data.hasOwnProperty('dim')) return 'dim'
    if(data.hasOwnProperty('cct')) return 'cct'
}