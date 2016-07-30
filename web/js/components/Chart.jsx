import React, {Component} from 'react'
var RadarChart = require('react-chartjs').Radar;

export default class Chart extends Component {
	render() {
        const chartData = {
            datasets: {
                0: {
                    data: [1,2,3,4,5,6],
                    fillColor: 'rgba(151,187,205,0.2)',
                    label: 'Policy Simulator Data',
                    pointColor: 'rgba(151,187,205,1)',
                    pointHighlightFill: '#fff',
                    pointHighlightStroke: 'rgba(220,220,220,1)',
                    pointStrokeColor: '#fff',
                    strokeColor: 'rgba(151,187,205,1)'
                }
            },
            labels: ['Culture', 'Apocalypse Readiness', 'Accessibility', 'Services', 'Safety', 'Community']
        };
		return (
            <div className='chart'>
                <RadarChart data={chartData} responsive/>
            </div>
        );
	}
}