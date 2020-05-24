import * as React from "react";
import Badge from "react-bootstrap/Badge";
import Container from "react-bootstrap/Container";
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";

export default class EndPointStatus extends React.Component {
    constructor(props) {
        super(props);
    }

    getBadge(endPoint) {
        if (endPoint.HelthStat.status)
            return <Badge variant="success">Success</Badge>
        return <Badge variant="danger">Fail</Badge>
    }

    getTimeTaken(endPoint){
        if(endPoint.HelthStat.timeTaken=="")
            return <span>{"NA"}</span>
        return <span>{endPoint.HelthStat.timeTaken}</span>
    }

    render() {
        const endPoint = this.props.endPoint;
        return <div>
                    <Container>
                        <Row>
                            <Col><span>{endPoint.ServiceInfo.ServiceName}</span></Col>
                            <Col>{this.getTimeTaken(endPoint)}</Col>
                            <Col>{this.getBadge(endPoint)}</Col>
                        </Row>
                    </Container>
                </div>
    }
}