import React, {Component} from 'react'
import './App.css'
import Form from './Form'
import InvoiceList from './InvoiceList'

class App extends Component {
  constructor(props) {
    super(props);

    this.state = {
      apiUrl: null,
      invoices: []
    }

    this.addInvoice = this.addInvoice.bind(this);
  }

  addInvoice(invoice) {
    let invoices = this.state.invoices;
    invoices.push(invoice);
    this.setState({invoices});
  }

  componentDidMount() {
    fetch('/api/url')
      .then( response => response.text())
      .then( url => {
        fetch(url + '/api/invoices')
          .then( response => response.json())
          .then( data => this.setState({invoices: data}));

        this.setState({apiUrl: url})
      });
  }

  render() {
    let { apiUrl } = this.state;
    if(apiUrl) {
      return (
        <div className="App container">
          <header className="App-header">
            <h1>E-invoice</h1>
          </header>
          <Form apiUrl={apiUrl} addInvoice={this.addInvoice} />
          <InvoiceList invoices={this.state.invoices} apiUrl={apiUrl}/>
        </div>
      );
    } else return <div></div>;
  }
}

export default App;
