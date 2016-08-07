/*
 * File: LinkList.java
 * Author: wdeqin
 * Time: 2015/5/23 1:21:39
 * Purpose: Java LinkList implement experiment
 */
class LinkList {

    private Node __head;
    private Node __tail;
    public LinkList() {
        this.__head = null;
        this.__tail = null;
    }

    public Node getHead() {
        return this.__head;
    }

    public Node getTail() {
        return this.__tail;
    }

    public void enqueue(Node node) {
        if (this.__tail == null) {
            this.__head = this.__tail = node;
            node.setNext(null);
        } else {
            this.__tail.setNext(node);
            this.__tail = node;
        }
    }

    public Node dequeue() {
        if (this.__head != null) {
            Node n = this.__head;
            this.__head = n.getNext();
            if (this.__head == null) {
                this.__tail = null;
            }
            n.setNext(null);
            return n;
        } else {
            return null;
        }
    }

    public boolean isEmpty() {
        return this.__head == null;
    }

}

class Node {
    private int value;
    private Node next;

    public Node(int value) {
        this.value = value;
        this.next = null;
    }

    public void setNext(Node next) {
        this.next = next;
    }

    public Node getNext() {
        return this.next;
    }

    public void setValue(int value) {
        this.value = value;
    }

    public int getValue() {
        return this.value;
    }
}

public class LinkListTest {
    public static void main(String[] argv) {
        LinkList l = new LinkList();
        for (int i = 0; i != 10; ++i) {
            Node n = new Node(i);
            l.enqueue(n);
        }

        System.out.println(l.dequeue().getValue());
        l.enqueue(new Node(19));
        System.out.println(l.dequeue().getValue());
        System.out.println(l.dequeue().getValue());
        l.enqueue(new Node(19));
        l.enqueue(new Node(19));
        System.out.println(l.dequeue().getValue());
        System.out.println(l.dequeue().getValue());

        while(!l.isEmpty()) {
            System.out.println(l.dequeue().getValue());
        }
        l.enqueue(new Node(19));
        System.out.println(l.dequeue().getValue());

    }
}

