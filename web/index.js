import App from './js/components/App';
import React from 'react'
import ReactDOM from 'react-dom'

require('file?name=[name].[ext]!./assets/index.html');

import './css/reset.css'
import './css/application.css'


var searchApp = ReactDOM.render(
	<App/>,
	document.getElementById('app')
)


function loaderMessagesIndex() {
    return Math.random() * loaderMessages.length;
}

function load() {
    var i = parseInt(loaderMessagesIndex(), 10);
    var loaderMessage = loaderMessages[i];
    searchApp.setState({loaderMessage: loaderMessage})
    var timer = setTimeout(function(){
        load();
    }, 200);
}
// load();
