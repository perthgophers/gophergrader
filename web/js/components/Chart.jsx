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
                    // rgb(86,128,181)
                    data: points,
                    backgroundColor: 'rgba(86,128,181,1)',
                    borderColor: 'rgba(179,181,198,1)',
                    pointBackgroundColor: 'rgba(179,181,198,1)',
                    pointBorderColor: 'rgba(86,128,181,1)',
                    pointHoverBackgroundColor: '#fff',
                    pointHoverBorderColor: 'rgba(179,181,198,1)',
                    label: 'Historical Data'
                }
            ],
            labels: labels
        };
            // defaultFontColor: '#666'
        const secondaryColor = "rgba(179,181,198,1)"
        const chartOptions = {
            scale: {
                angleLines: {
                    color: secondaryColor,
                    lineWidth: 2
                },
                gridLines: {
                    color: secondaryColor,
                    lineWidth: 2
                },
                ticks: {
                    color: secondaryColor,
                    fixedStepSize: 2,
                    suggestedMax: 10,
                    backdropColor: 'rgba(0,0,0,0)'
                },
                elements: {
                    line: {
                        lineTension: 1000
                    }
                },
                pointLabels: {
                    fontColor: secondaryColor,
                    fontSize: 14
                }
            },
        }

		return (
            <div className='chart'>
                <Radar data={chartData} options={chartOptions}/>
            </div>
        );
	}
}