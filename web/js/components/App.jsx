import React, {Component} from 'react'
import Chart from './Chart'
import * as axios from 'axios';

export default class App extends Component {
    constructor(props) {
        super(props)
        this.state = {
            points: [0, 0, 0, 0, 0, 0],
            labels: ['Culture', 'Apocalypse Readiness', 'Accessibility', 'Services', 'Safety', 'Community'],
            searching: false,
            loaderMessage: '', 
        };
        this.search = this.search.bind(this);
    }

    search() {
        console.log('loading');
        axios.get("http://localhost:9000/grade?address=82 East St, Maylands, WA, 6051")
            .then((response) => {
                console.log('load complete');
                const currentState = this.state;
                const nextState = currentState;
                
                nextState.points = [
                    Math.floor(Math.random()*10),
                    Math.floor(Math.random()*10),
                    Math.floor(Math.random()*10),
                    Math.floor(Math.random()*10),
                    Math.floor(Math.random()*10),
                    Math.floor(Math.random()*10)
                ];
                console.log(nextState)
                this.setState(nextState);
            })
            .catch((err) => {
                console.log(`Error: ${err}`);
            });
    }

	render() {
		return <div>
            <div className="search-box">
                <input className="search-input" />
                <button className="search-button" onClick={this.search}>Search</button>
            </div>
            <div className="search-result">
                <div className="loader blink">
                    {this.state.loaderMessage}
                </div>
            </div>
            <Chart points={this.state.points} labels={this.state.labels}/>
        </div>
	}
}