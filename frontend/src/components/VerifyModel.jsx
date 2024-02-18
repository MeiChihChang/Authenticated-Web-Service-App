import {Modal, Input, Form, Button, ModalHeader, ModalBody, ModalFooter} from 'reactstrap'

/**
 * @description This component renders a verifier model dialog to the user.
 *
 * @param {boolean} isopen keep this model dialog open or hidden.
 * @param {string} test the test string.
 * @param {function} handler a function to handle onSubmit event.
 * @param {string} error error string.
 * @returns {VerifyModal} A React element that renders a verifier model dialog to ask the user to pass a mathematic test.
 */
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
