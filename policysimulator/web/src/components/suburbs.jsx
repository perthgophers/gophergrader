import React, {Component} from 'react';

import {Table, Column, Cell} from 'fixed-data-table';
import superagent from 'superagent'

//A rowGetter function is required by the grid to retrieve a row for a given index

export default class SuburbsTable extends Component{
  constructor(props) {
    super(props)
    this.state = {suburbs:[]}
    superagent.get('/suburbs')
      .end((err, res) => {
        if (err) {
          alert(err);
          return
        };
        this.setState({suburbs: JSON.parse(res.text)})
      }
    );
  }
	render() {
    let rows = this.state.suburbs
    console.log("ROWS,",rows[0])
		return <div className="suburbs">
      <Table
          rowHeight={50}
          rowsCount={rows.length}
          width={5000}
          height={5000}
          headerHeight={50}>
          <Column
            header={<Cell>Suburbs</Cell>}
            cell={({rowIndex, ...props}) => (
              <Cell {...props}>
                <a onClick={this.props.onClick}>{rows[rowIndex].name}</a>
              </Cell>
            )}
            width={2000}
          />
        </Table>
		</div>
	}
}