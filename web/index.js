import App from './js/components/App';

require('file?name=[name].[ext]!./assets/index.html');

import './css/reset.css'
import './css/application.css'

let app = new App()

app.init()