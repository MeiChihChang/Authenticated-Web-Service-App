import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import { Container, Row, Col} from 'reactstrap';

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <Container className="mt-5"> 
      <Row className="justify-content-center"> 
      <Col xs={12} md={8}>
        <App />
      </Col>
      </Row>
    </Container>
  </React.StrictMode>
);

