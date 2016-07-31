
import GoldenLayout from './goldenlayout'

import {On, Off, Emit} from './events.js'
import React, {Component} from 'react'
import ReactDOM from 'react-dom';

import Immutable from 'immutable'
import superagent from 'superagent'
import SuburbsTable from './suburbs.jsx'
import FormsComponet from './form.jsx'
import GetStore from './store.js'

import Chart from './charts.jsx'

import * as axios from 'axios';

class Blank extends Component {
	render() {
		return <div />
	}
}

function randNum() {
	var min = 1;
	var max = 10;
	return Math.floor(Math.random() * (max - min + 1)) + min;
}

class ChartsComponent extends Component {
	search = (event) => {
	  this.setState({searching: true});
	  const query = "Maylands";
	  const fullQuery = `/suburb?address=${query}`
	  axios.get(fullQuery)
	      .then((response) => {
	          console.log('load complete');
	          console.log(response);
	          const currentState = this.state;
	          const nextState = currentState;
	          nextState.points = [
	              response.data.accessibility,
	              response.data.apocalypse,
	              response.data.community,
	              response.data.culture,
	              response.data.safety,
	              response.data.services
	          ];

	          nextState.searchComplete = true;
	          nextState.searching = false;
	          this.setState(nextState);
	      })
	      .catch((err) => {
	          console.log(`Error: ${err}`);
	          this.setState({searching: false});
	      });
	}
	constructor(props) {
		super(props)
		this.state = {
      searchText: '',
      points: [0, 0, 0, 0, 0, 0],
      points2: [randNum(), randNum(), randNum(), randNum(), randNum(), randNum()],
      labels: [
          'Accessibility', 
          'Apocalypse Readiness', 
          'Community',
          'Culture', 
          'Safety', 
          'Services' 
       ],
      searching: false,
      searchComplete: false,
      loaderMessage: '', 
      errorMessage: ''
  	};
  	On('update',() => {
  		this.setState({
  			points2: [randNum(), randNum(), randNum(), randNum(), randNum(), randNum()]
  		})
  	})
	}
	componentDidMount() {
	 	this.search()     
	}
	renderChart = () => {
	  if (this.state.searchComplete) {
	      return <Chart points2={this.state.points2} points={this.state.points} labels={this.state.labels}/> 
	  }
	}
	render() {
		return (<div>
				{this.renderChart()}
			</div>
		)
	}
}

export default class App {
	loadSuburb = () => {

	}
	init = () => {
		this.store = GetStore()
		this.layout = new GoldenLayout({
		 settings: {
        selectionEnabled: true
     },
		 content: [{
			 type: 'row',
			 content:[{
				 type:'react-component',
				 component: 'suburbs',
				 width: 25,
				 title: "Suburbs",
				 isCloseable: false,
				 props: { label: 'Suburbs', onClick:this.loadSuburb, store:this.store }
			 },{
				 type:'react-component',
				 component: 'forms',
				 width: 22,
				 title: "Functions",
				 isCloseable: false,
				 props: { label: 'Suburbs', store:this.store }
			 	},{
				 type:'react-component',
				 component: 'chart',
				 width: 60,
				 title: "Visualiser",
				 isCloseable: false,
				 props: { label: 'Suburbs', store:this.store }
			}]
		 }]
		})
		this.layout.registerComponent( 'suburbs', SuburbsTable );
		this.layout.registerComponent( 'chart', ChartsComponent );
		this.layout.registerComponent( 'forms', FormsComponet );
		this.layout.registerComponent( 'blank', Blank );

		this.layout.init();

	}
}

