import React, {Component} from 'react'
import Chart from './Chart'
import LoadingText from './LoadingText';

import * as axios from 'axios';

export default class App extends Component {
    constructor(props) {
        super(props)
        const loaderMessages = [
            'Adding Hidden Agendas ',
            'Adjusting Bell Curves ',
            'Aesthesizing Industrial Areas ',
            'Aligning Covariance Matrices ',
            'Applying Feng Shui Shaders ',
            'Applying Theatre Soda Layer ',
            'Asserting Packed Exemplars ',
            'Attempting to Lock Back-Buffer ',
            'Binding Sapling Root System ',
            'Breeding Fauna ',
            'Building Data Trees ',
            'Bureacritizing Bureaucracies ',
            'Calculating Inverse Probability Matrices ',
            'Calculating Llama Expectoration Trajectory ',
            'Calibrating Blue Skies ',
            'Charging Ozone Layer ',
            'Coalescing Cloud Formations ',
            'Cohorting Exemplars ',
            'Collecting Meteor Particles ',
            'Compounding Inert Tessellations ',
            'Compressing Fish Files ',
            'Computing Optimal Bin Packing ',
            'Concatenating Sub-Contractors ',
            'Containing Existential Buffer ',
            'Debarking Ark Ramp ',
            'Debunching Unionized Commercial Services ',
            'Deciding What Message to Display Next ',
            'Decomposing Singular Values ',
            'Decrementing Tectonic Plates ',
            'Deleting Ferry Routes ',
            'Depixelating Inner Mountain Surface Back Faces ',
            'Depositing Slush Funds ',
            'Destabilizing Economic Indicators ',
            'Determining Width of Blast Fronts ',
            'Deunionizing Bulldozers ',
            'Dicing Models ',
            'Diluting Livestock Nutrition Variables ',
            'Downloading Satellite Terrain Data ',
            'Exposing Flash Variables to Streak System ',
            'Extracting Resources ',
            'Factoring Pay Scale ',
            'Fixing Election Outcome Matrix ',
            'Flood-Filling Ground Water ',
            'Flushing Pipe Network ',
            'Gathering Particle Sources ',
            'Generating Jobs ',
            'Gesticulating Mimes ',
            'Graphing Whale Migration ',
            'Hiding Willio Webnet Mask ',
            'Implementing Impeachment Routine ',
            'Increasing Accuracy of RCI Simulators ',
            'Increasing Magmafacation ',
            'Initializing My Sim Tracking Mechanism ',
            'Initializing Rhinoceros Breeding Timetable ',
            'Initializing Robotic Click-Path AI ',
            'Inserting Sublimated Messages ',
            'Integrating Curves ',
            'Integrating Illumination Form Factors ',
            'Integrating Population Graphs ',
            'Iterating Cellular Automata ',
            'Lecturing Errant Subsystems ',
            'Mixing Genetic Pool ',
            'Modeling Object Components ',
            'Mopping Occupant Leaks ',
            'Normalizing Power ',
            'Obfuscating Quigley Matrix ',
            'Overconstraining Dirty Industry Calculations ',
            'Partitioning City Grid Singularities ',
            'Perturbing Matrices ',
            'Pixalating Nude Patch ',
            'Polishing Water Highlights ',
            'Populating Lot Templates ',
            'Preparing Sprites for Random Walks ',
            'Prioritizing Landmarks ',
            'Projecting Law Enforcement Pastry Intake ',
            'Realigning Alternate Time Frames ',
            'Reconfiguring User Mental Processes ',
            'Relaxing Splines ',
            'Removing Road Network Speed Bumps ',
            'Removing Texture Gradients ',
            'Removing Vehicle Avoidance Behavior ',
            'Resolving GUID Conflict ',
            'Reticulating Splines ',
            'Retracting Phong Shader ',
            'Retrieving from Back Store ',
            'Reverse Engineering Image Consultant ',
            'Routing Neural Network Infrastructure ',
            'Scattering Rhino Food Sources ',
            'Scrubbing Terrain ',
            'Searching for Llamas ',
            'Seeding Architecture Simulation Parameters ',
            'Sequencing Particles ',
            'Setting Advisor Moods ',
            'Setting Inner Deity Indicators ',
            'Setting Universal Physical Constants ',
            'Sonically Enhancing Occupant-Free Timber ',
            'Speculating Stock Market Indices ',
            'Splatting Transforms ',
            'Stratifying Ground Layers ',
            'Sub-Sampling Water Data ',
            'Synthesizing Gravity ',
            'Synthesizing Wavelets ',
            'Time-Compressing Simulator Clock ',
            'Unable to Reveal Current Activity ',
            'Weathering Buildings ',
            'Zeroing Crime Network',
        ];
        this.state = {
            loaderMessages: loaderMessages,
            searchText: '',
            points: [0, 0, 0, 0, 0, 0],
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
        this.search = this.search.bind(this);
        this.renderChart = this.renderChart.bind(this);
        this.handleChange = this.handleChange.bind(this);
        this.handleEnter = this.handleEnter.bind(this);
    }



    search(event) {
        this.setState({searching: true});
        const query = this.state.searchText;
        const partialQuery = `/grade?address=${query}`;
        const fullQuery = `http://localhost:9000/grade?address=${query}`
        axios.get(partialQuery)
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
    handleChange(event) {
        this.setState({searchText: event.target.value});
    }

    handleEnter(event) {
        if (event.key === 'Enter') {
            this.search()
        }
    }

	render() {
        const chart = this.renderChart();
        const loader = this.renderLoader();
        const searching = this.state.searching;
        const searchComplete = this.state.searchComplete;
        let style = {};
        if (searching || searchComplete) {
            style = {
                top: 30,
                transform: 'translate(-50%)'
            }
        }

            
		return <div className="search-container" style={style}>
            <div className="logo" />
            <div className='search-box'>
                <input className='search-input' onChange={this.handleChange} onKeyDown={this.handleEnter}/>
                <button className='search-button' onClick={this.search}>Search</button>
            </div>
            {loader}
            {chart}
        </div>
	}

    renderLoader() {
        if (this.state.searching) {
            return <div className='search-result'>
                <div className='loader blink'>
                    <LoadingText searching={this.state.searching} loaderMessages={this.state.loaderMessages}/>
                </div>
            </div>
        }
    }

    renderChart() {
        if (this.state.searchComplete && !this.state.searching) {
            return <Chart points={this.state.points} labels={this.state.labels}/> 
        }
    }
}