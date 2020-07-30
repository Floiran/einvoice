import React, {Component} from "react";

function Invoice(props) {
    return <div>
        <p>Id: {props.id}</p>
        <p>Sender: {props.sender}</p>
        <p>Receiver: {props.receiver}</p>
    </div>
}

class Form extends Component {
    constructor(props) {
        super(props);

        this.state = {
            getInvoice: null,
            postInvoice: null,
            getInvoiceId: "",
            postInvoiceSender: "",
            postInvoiceReceiver: ""
        }

        this.handleInputChange = this.handleInputChange.bind(this);
        this.submitGetInvoice = this.submitGetInvoice.bind(this);
        this.submitPostInvoice = this.submitPostInvoice.bind(this);
    }

    handleInputChange(event) {
        const target = event.target;
        const name = target.name;
        this.setState({
            [name]: target.value
        });
    }

    submitGetInvoice() {
        console.log("get");
        fetch('/invoice/' + this.state.getInvoiceId)
            .then( response => response.json())
            .then( data => this.setState({getInvoice: data}));
    }

    submitPostInvoice() {
        fetch('/invoice' + this.state.getInvoiceId, {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                sender: this.state.postInvoiceSender,
                receiver: this.state.postInvoiceReceiver
            })
        })
            .then( response => response.json())
            .then( data => this.setState({postInvoice: data}));
    }

    render() {
        let { getInvoice, postInvoice } = this.state;

        return (
            <div>
                <h2>Get invoice</h2>
                <label>
                    Invoice id:
                    <input type="text" name="getInvoiceId" value={this.state.getInvoiceId} onChange={this.handleInputChange} />
                </label>
                <button className='btn btn-primary' onClick={this.submitGetInvoice}>Submit</button>
                {getInvoice && <Invoice {...getInvoice} />}
                <h2>Post invoice</h2>
                <p>
                    Sender:
                    <input type="text" name="postInvoiceSender" value={this.state.postInvoiceSender} onChange={this.handleInputChange} />
                </p>
                <p>
                    Receiver:
                    <input type="text" name="postInvoiceReceiver" value={this.state.postInvoiceReceiver} onChange={this.handleInputChange} />
                </p>
                <button className='btn btn-primary' onClick={this.submitPostInvoice}>Submit</button>
                {postInvoice && <Invoice {...postInvoice} />}
            </div>
        )
    }
}

export default Form;