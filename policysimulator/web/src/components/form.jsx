import React, {Component} from 'react';
import InputSlider from 'react-input-slider';

import {On, Off, Emit} from './events.js'

import 'react-input-slider/dist/input-slider.css'

export default React.createClass({
  getInitialState() {
    return {
      x: 21,
      y: 73,
      left: 120,
      top: 120,
      pstations: 1,
      hospitals: 0,
      health: 2,
      worship: 4,
      schools: 1
    };
  },
  timer: null,
  componentDidUpdate() {
    clearTimeout(this.timer)
    this.timer = setTimeout(() => {
      Emit('update')
    },500)
  },

  render() {
    return (
      <div className="form">
        <div className="example example-x">
          <h2>{'Police Stations: ' + this.state.pstations}</h2>
          <div></div>
          <InputSlider
            className="slider"
            axis="x"
            x={this.state.pstations}
            xmin={0}
            xmax={5}
            onDragEnd={this.handleDragEnd}
            onChange={(pos) => {
              this.setState({
                pstations: Math.round(pos.x)
              });
            }}
          />
        </div>
        <div className="example example-x">
          <h2>{'Hospitals: ' + this.state.hospitals}</h2>
          <div></div>
          <InputSlider
            className="slider"
            axis="x"
            x={this.state.hospitals}
            xmin={0}
            xmax={5}
            onDragEnd={this.handleDragEnd}
            onChange={(pos) => {
              this.setState({
                hospitals: Math.round(pos.x)
              });
            }}
          />
        </div>
        <div className="example example-x">
          <h2>{'Health care facilities: ' + this.state.health}</h2>
          <div></div>
          <InputSlider
            className="slider"
            axis="x"
            x={this.state.health}
            xmin={0}
            xmax={5}
            onDragEnd={this.handleDragEnd}
            onChange={(pos) => {
              this.setState({
                health: Math.round(pos.x)
              });
            }}
          />
        </div>
        <div className="example example-x">
          <h2>{'Places of worship: ' + this.state.worship}</h2>
          <div></div>
          <InputSlider
            className="slider"
            axis="x"
            x={this.state.worship}
            xmin={0}
            xmax={5}
            onDragEnd={this.handleDragEnd}
            onChange={(pos) => {
              this.setState({
                worship: Math.round(pos.x)
              });
            }}
          />
        </div>
        <div className="example example-x">
          <h2>{'Schools: ' + this.state.schools}</h2>
          <div></div>
          <InputSlider
            className="slider"
            axis="x"
            x={this.state.schools}
            xmin={0}
            xmax={5}
            onDragEnd={this.handleDragEnd}
            onChange={(pos) => {
              this.setState({
                schools: Math.round(pos.x)
              });
            }}
          />
        </div>

        <h2>{'Hospitals: ' + this.state.left}</h2>
        <div className="example example-x">
          <div></div>
          <InputSlider
            className="slider"
            axis="x"
            x={this.state.left}
            xmin={100}
            xmax={360}
            onDragEnd={this.handleDragEnd}
            onChange={this.handleChangeX}
          />
        </div>
      </div>
    );
  },

  handleDragEnd() {
    console.log('handleDragEnd');
  },

  handleChange(pos) {
    this.setState({
      x: pos.x,
      y: pos.y
    });
  },

  

});
