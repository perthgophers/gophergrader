import React, {Component} from 'react'

export default class LoadingText extends Component {
    constructor(props) {
        super(props);
        this.loadText = this.loadText.bind(this);
    }
	render() {
        const searching = this.props.searching;
        const message = this.loadText();
        console.log(message)
        if (searching) {
            return (
                <div className='search-result'>
                    Loading
                    {message}
                </div>
            );
        } else {
            return(<div/>);
        }
	}

    loadText() {
        console.log('sampler')
        const loaderMessages = this.props.loaderMessages;
        var i = parseInt(() => {
            Math.random() * loaderMessages.length;
        }, 10);
        var loaderMessage = loaderMessages[i];
        console.log(loaderMessage)
        return loaderMessage
    }
}

