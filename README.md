# Design patterns in golang

## Strategy
Define a family of algorithms, encapsulate each one, and make them interchangeable. Strategy lets the algorithm vary independently from clients that use it.

## Singleton
Ensure a class only has one instance, and provide a global point of access to it.

## Composite
Composite is a structural design pattern that lets you compose objects into tree structures and then work with these structures as if they were individual objects.

 ### Pros
 #### You can work with complex tree structures more conveniently: use polymorphism and recursion to your advantage.
 ####  Open/Closed Principle. You can introduce new element types into the app without breaking the existing code, which now works with the object tree.
 ### Cons
 #### It might be difficult to provide a common interface for classes whose functionality differs too much. In certain scenarios, you’d need to overgeneralize the component interface, making it harder to comprehend.