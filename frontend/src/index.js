import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";

class FieldInput extends React.Component {
  constructor() {
    super();
    this.state = {
      data: [],
    };
  }

  handleInput = (event) => {
    fetch(`http://localhost:8000/api?q=${event.target.value}`, {
      method: "GET",
    })
      .then((res) => res.json())
      .then((res) => {
        if (!event.target.value.length > 0) {
          res = [];
        }

        this.setState({
          data: res,
        });
      });
  };

  handleClick = (event) => {
    console.log(event.target.value);

    this.setState({
      input: this.state.data[event.target.value].password,
    });
  };

  render() {
    var data = this.state.data || [];
    var input = this.state.input;
    let final = [];

    for (let [i, elem] of data.entries()) {
      final.push(
        <button value={i} onClick={this.handleClick} key={i}>
          {elem.password}
        </button>
      );
    }

    return (
      <div className="container-field-input">
        <h1 className="page-title">Pick a password that suits best for YOU!</h1>
        <input onChange={this.handleInput} value={input}></input>
        <div className="suggestions">{final}</div>
      </div>
    );
  }
}

// ========================================

const root = ReactDOM.createRoot(document.getElementById("root"));
root.render(<FieldInput />);
