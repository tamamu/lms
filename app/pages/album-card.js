
import React from 'react'

const cardStyle = {
  backgroundColor: "#37474F",
  color: "white",
  width: 240,
  height: 160,
  border: "1px solid black"
};

const titleStyle = {
  fontSize: "1.2em"
};

const propStyle = {
  display: "inline",
  marginRight: "1em"
};

export default class AlbumCard extends React.Component {
  render() {
    return (
      <div album-id={this.props.albumId} artist-id={this.props.artistId} key={this.props.key} style={cardStyle}>
        <p style={titleStyle}>{this.props.title}</p>
        <p style={propStyle}>{this.props.genre}</p>
        <p style={propStyle}>{this.props.year}</p>
      </div>
    )
  }
}

/*
AlbumCard.propTypes = {
  title: React.PropTypes.string.isRequired,
  genre: React.PropTypes.string,
  year: React.PropTypes.number
}
*/
