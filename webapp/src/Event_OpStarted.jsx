import React from 'react';

export default function EventOpStarted(props) {
  return (
    <div style={{color: 'rgb(96, 253, 255)'}}>
      OpStarted
      Id='{props.opStarted.opId}'
      PkgRef='{props.opStarted.pkgRef}'
      Timestamp='{props.timestamp}'
    </div>
  );
}
