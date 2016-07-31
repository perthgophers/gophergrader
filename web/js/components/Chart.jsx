import React, {Component} from 'react'
import {Radar} from 'react-chartjs-2';

export default class Chart extends Component {
    constructor(props) {
        super(props);
    }
	render() {

        const points = this.props.points;
        const labels = this.props.labels;
        const chartData = {
            datasets: [
                {
                    data: points,
                    backgroundColor: 'rgba(179,181,198,0.2)',
                    borderColor: 'rgba(179,181,198,1)',
                    pointBackgroundColor: 'rgba(179,181,198,1)',
                    pointBorderColor: '#fff',
                    pointHoverBackgroundColor: '#fff',
                    pointHoverBorderColor: 'rgba(179,181,198,1)'

                }
            ],
            labels: labels
        };
            // defaultFontColor: '#666'
        const chartOptions = {
        }

		return (
            <div className='chart'>
                <Radar data={chartData}/>
            </div>
        );
	}
}