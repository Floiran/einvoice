import React, {Component} from 'react';
import './App.css';
import CreateInvoice from './CreateInvoice';
import InvoiceList from './InvoiceList';

class App extends Component {
  constructor(props) {
    super(props);

    this.state = {
      apiUrl: null,
      invoices: [],
      createInvoice: false
    }

    this.addInvoice = this.addInvoice.bind(this);
    this.createInvoice = this.createInvoice.bind(this);
  }

  addInvoice(invoice) {
    let invoices = this.state.invoices;
    invoices.push(invoice);
    this.setState({invoices});
  }

  createInvoice(){
    this.setState({ createInvoice: !this.state.createInvoice});
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
          <header className="App-header row">
            <h1 className='col'>E-invoice</h1>
          </header>
          <div className='row justify-content-center'>
            <button className='btn btn-primary' onClick={this.createInvoice} >Create invoice</button>
          </div>
          { this.state.createInvoice && <CreateInvoice apiUrl={apiUrl} addInvoice={this.addInvoice} /> }
          <InvoiceList invoices={this.state.invoices} apiUrl={apiUrl}/>
        </div>
      );
    } else return <div></div>;
  }
}

export default App;
