// https://reactjs.org/docs/components-and-props.html 
import React from 'react'
import ReactDOM from 'react-dom'

//function classNameDinamic() {
    //return "kucing" + (5+5)
//}
//function damnElement() {
    //return <h2 className={classNameDinamic()}>makan nasi: <strong>It is {new Date().toLocaleTimeString()} </strong>  </h2>
//}

//function tick(){
    //ReactDOM.render(
        //damnElement(), 
        //document.getElementById('root')
    //);
//}

/*setInterval(tick, 1000 );*/

class Clock extends React.Component {
    constructor(props) {
        super(props);
        this.state = {date: new Date()};
    }

    // lifecycle hooks
    componentDidMount() {
        this.waka = setInterval(
            () => this.tick(),
            1000
        ); 
    }

    componenWillMount() {
        clearInterval(this.waka); 
    }

    tick() {
        this.setState({
            date: new Date()
        });
    }

    render() {
        return (
            <div>
                <h1>Hello, world!</h1>
                <h2>It is {this.state.date.toLocaleTimeString()}.</h2>
            </div>
        );
    }
}

ReactDOM.render(
  <Clock />,
  document.getElementById('root')
);

// ---------------------------------------------------------------------

var elament1 = React.createElement(
    'h3',
    {className: 'Greeting'+(3+3)},
    'hallo brader'
);

ReactDOM.render(
    elament1, 
    document.getElementById('root1')
);

// ---------------------------------------------------------------------

// Component

function Welcome(props) {
  return <h1>Hello, {props.name} from {props.from}</h1>;
}

var App = function() {
    return (
        <div>
            <Welcome name="Sara oi" from="bandung"/>
            <Welcome name="Jaja" from="bandung"/>
            <Welcome name="Dora" from="bandung"/>
        </div>
    )
};





function Avatar(props) {
  return (
    <img className="Avatar"
      src={props.user.avatarUrl}
      alt={props.user.name}
    />
  );
}

function UserInfo(props) {
    return (
    <div className="UserInfo">
          <Avatar user={props.author} /> 
          <div className="UserInfo-name">
          Pakde: {props.author.name}
        </div>
    </div>
    );
} 

function CommentDate(props) {
    return (
        <div className="Comment-date">
            {props.date.toLocaleTimeString()}
        </div>
    )
}

function CommentText (props) {
    return (
        <div className="Comment-text">
            {props.text}
        </div>
    )
} 
function Comment(props) {
  return (
    <div className="Comment">
        <UserInfo author={props.author} />
        <CommentDate date={props.date} />
        <CommentText text={props.text}/>
    </div>
  );
}



const comment = {
  date: new Date(),
  text: 'I hope you enjoy learning React!',
  author: {
    name: 'Hello Kitty',
    avatarUrl: 'http://placekitten.com/g/64/64'
  }
};

ReactDOM.render(
  <Comment author={comment.author} date={comment.date} text={"keren sekali"} />,
  document.getElementById('root2')
);

