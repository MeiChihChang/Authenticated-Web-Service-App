import {Modal, Input, Form, Button, ModalHeader, ModalBody, ModalFooter} from 'reactstrap'

function VerifyModal({ isopen, test, handler, error }) {
    return (
        <Modal isOpen={isopen}>
            <ModalHeader>Question:</ModalHeader>
            <Form onSubmit={handler}>
            <ModalBody>
                <p>
                    What is the result of this <strong>{test} ?</strong>
                </p>
                <h5>Answear:</h5>
                
                    <Input
                        type="number"
                        id="answear"
                        name="answear"
                        required
                    />     
                    
                
            </ModalBody>
            <ModalFooter>
            {error !== "" && <p style={{ color: "red" }}>{error}</p>} 
            <Button >Submit</Button>
            </ModalFooter>
            </Form>
        </Modal>
    );
}

export default VerifyModal;
