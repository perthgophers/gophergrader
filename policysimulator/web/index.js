import App from './src/components/app';
import React from 'react'
import ReactDOM from 'react-dom'

require('file?name=[name].[ext]!./src/assets/index.html');


import './src/css/reset.css'
import './src/css/goldenlayout-base.css'
import './src/css/goldenlayout-dark.css'
import './src/css/application.css'
import './src/css/fixed-data-table.css'

let psapp = new App()

psapp.init()