import React, {Component} from 'react';
import './App.css';
import Form from "./Form";

class App extends Component {
    constructor(props) {
        super(props);

        this.state = {
            apiUrl: null
        }
    }

    componentDidMount() {
        fetch('/api/url')
            .then( response => response.text())
            .then( data => this.setState({apiUrl: data}));
    }

    render() {
        let { apiUrl } = this.state;
        if(apiUrl) {
            return (
                <div className="App">
                    <header className="App-header">
                        <h1>E-invoice</h1>
                    </header>
                    <Form apiUrl={apiUrl} />
                </div>
            );
        } else return <div></div>;
    }

}

export default App;
