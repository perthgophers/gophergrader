import App from './js/components/App';
import React from 'react'
import ReactDOM from 'react-dom'

require('file?name=[name].[ext]!./assets/index.html');

import './css/reset.css'
import './css/application.css'


ReactDOM.render(
	<App/>,
	document.getElementById('app')
)