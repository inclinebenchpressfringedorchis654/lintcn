# tsgolint Rules Reference

All 59 type-checked rules implemented by tsgolint, with wrong/right code examples.
Rules marked with **fix** have auto-fix or suggestion support.

---

## 1. await-thenable **fix**

Disallow awaiting a value that is not a Thenable.

```ts
// Wrong
await 'value';
```

```ts
// Right (fixed: remove await)
'value';
```

---

## 2. consistent-return

Require return statements to either always or never specify values.

```ts
// Wrong
function foo(flag: boolean): undefined {
  if (flag) return 'yes';
  return;
}
```

```ts
// Right
function foo(flag: boolean): string {
  if (flag) return 'yes';
  return 'no';
}
```

---

## 3. consistent-type-exports **fix**

Enforce consistent usage of type exports.

```ts
// Wrong
interface ButtonProps { onClick: () => void }
export { ButtonProps };
```

```ts
// Right (fixed: add type keyword)
export type { ButtonProps };
```

---

## 4. dot-notation **fix**

Enforce dot notation whenever possible.

```ts
// Wrong
declare const obj: { name: string };
const val = obj['name'];
```

```ts
// Right (fixed: use dot notation)
declare const obj: { name: string };
const val = obj.name;
```

---

## 5. no-array-delete **fix**

Disallow using the `delete` operator on array values.

```ts
// Wrong
declare const arr: number[];
delete arr[0];
```

```ts
// Right (fixed: use splice)
declare const arr: number[];
arr.splice(0, 1);
```

---

## 6. no-base-to-string

Require `.toString()` to only be called on objects which provide useful information when stringified.

```ts
// Wrong
class Foo {}
const foo = new Foo();
const str = `Value: ${foo}`;
```

```ts
// Right
class Foo { toString() { return 'Foo'; } }
const foo = new Foo();
const str = `Value: ${foo}`;
```

---

## 7. no-confusing-void-expression **fix**

Require expressions of type void to appear in statement position.

```ts
// Wrong
const response = alert('Are you sure?');
console.log(alert('click'));
```

```ts
// Right (fixed: separate into statement)
alert('Are you sure?');
```

---

## 8. no-deprecated

Disallow using code marked as `@deprecated`.

```ts
// Wrong
/** @deprecated Use apiV2 instead. */
declare function apiV1(): void;
apiV1();
```

```ts
// Right
declare function apiV2(): void;
apiV2();
```

---

## 9. no-duplicate-type-constituents **fix**

Disallow duplicate constituents of union or intersection types.

```ts
// Wrong
type T = string | string | number;
```

```ts
// Right (fixed: remove duplicate)
type T = string | number;
```

---

## 10. no-floating-promises **fix**

Require Promise-like statements to be handled appropriately.

```ts
// Wrong
async function fetchData() { return 'value'; }
fetchData();
```

```ts
// Right (fixed: add await or void)
await fetchData();
// or
void fetchData();
```

---

## 11. no-for-in-array

Disallow iterating over an array with a for-in loop.

```ts
// Wrong
declare const arr: string[];
for (const i in arr) { console.log(arr[i]); }
```

```ts
// Right
declare const arr: string[];
for (const value of arr) { console.log(value); }
```

---

## 12. no-implied-eval

Disallow the use of `eval()`-like functions.

```ts
// Wrong
setTimeout('alert("Hi!")', 100);
```

```ts
// Right
setTimeout(() => alert('Hi!'), 100);
```

---

## 13. no-meaningless-void-operator **fix**

Disallow the `void` operator except when used to discard a value.

```ts
// Wrong
function foo() {}
void foo();
```

```ts
// Right (fixed: remove void)
function foo() {}
foo();
```

---

## 14. no-misused-promises

Disallow Promises in places not designed to handle them.

```ts
// Wrong
const promise = Promise.resolve('value');
if (promise) { /* always truthy! */ }
```

```ts
// Right
const promise = Promise.resolve('value');
if (await promise) { /* ... */ }
```

---

## 15. no-misused-spread **fix**

Disallow using the spread operator when it might cause unexpected behavior.

```ts
// Wrong
declare const promise: Promise<number>;
const obj = { ...promise };
```

```ts
// Right (fixed: await first)
declare const promise: Promise<number>;
const obj = { ...(await promise) };
```

---

## 16. no-mixed-enums

Disallow enums from having both number and string members.

```ts
// Wrong
enum Status {
  Unknown,
  Closed = 1,
  Open = 'open',
}
```

```ts
// Right
enum Status {
  Unknown = 0,
  Closed = 1,
  Open = 2,
}
```

---

## 17. no-redundant-type-constituents

Disallow members of unions and intersections that do nothing or override type information.

```ts
// Wrong
type T = any | 'foo';
type U = string | 'literal';
```

```ts
// Right
type T = any;
type U = string;
```

---

## 18. no-unnecessary-boolean-literal-compare **fix**

Disallow unnecessary equality comparisons against boolean literals.

```ts
// Wrong
declare const isReady: boolean;
if (isReady === true) {}
```

```ts
// Right (fixed: remove literal compare)
declare const isReady: boolean;
if (isReady) {}
```

---

## 19. no-unnecessary-condition **fix**

Disallow conditionals where the type is always truthy or always falsy.

```ts
// Wrong
function head(items: string[]) {
  if (items) { return items[0]; }
}
```

```ts
// Right
function head(items: string[]) {
  if (items.length) { return items[0]; }
}
```

---

## 20. no-unnecessary-qualifier **fix**

Disallow unnecessary namespace qualifiers.

```ts
// Wrong
enum A {
  B,
  C = A.B,
}
```

```ts
// Right (fixed: remove qualifier)
enum A {
  B,
  C = B,
}
```

---

## 21. no-unnecessary-template-expression **fix**

Disallow unnecessary template expressions.

```ts
// Wrong
const text = 'hello';
const wrapped = `${text}`;
```

```ts
// Right (fixed: unwrap)
const text = 'hello';
const wrapped = text;
```

---

## 22. no-unnecessary-type-arguments **fix**

Disallow type arguments that are equal to the default.

```ts
// Wrong
function f<T = number>() {}
f<number>();
```

```ts
// Right (fixed: remove type arg)
function f<T = number>() {}
f();
```

---

## 23. no-unnecessary-type-assertion **fix**

Disallow type assertions that do not change the type of an expression.

```ts
// Wrong
const foo = 3;
const bar = foo!;
```

```ts
// Right (fixed: remove assertion)
const foo = 3;
const bar = foo;
```

---

## 24. no-unnecessary-type-conversion **fix**

Disallow conversion idioms when they do not change the type or value.

```ts
// Wrong
String('hello');
!!true;
+42;
```

```ts
// Right (fixed: remove conversion)
'hello';
true;
42;
```

---

## 25. no-unnecessary-type-parameters **fix**

Disallow type parameters that aren't used multiple times.

```ts
// Wrong
function parseJSON<T>(input: string): T {
  return JSON.parse(input);
}
```

```ts
// Right (fixed: use unknown)
function parseJSON(input: string): unknown {
  return JSON.parse(input);
}
```

---

## 26. no-unsafe-argument

Disallow calling a function with a value with type `any`.

```ts
// Wrong
declare function foo(arg: string): void;
const value = 1 as any;
foo(value);
```

```ts
// Right
declare function foo(arg: string): void;
foo('hello');
```

---

## 27. no-unsafe-assignment

Disallow assigning a value with type `any` to variables and properties.

```ts
// Wrong
const x = 1 as any;
const y: string[] = new Array<any>();
```

```ts
// Right
const x = 1;
const y: string[] = new Array<string>();
```

---

## 28. no-unsafe-call

Disallow calling a value with type `any`.

```ts
// Wrong
declare const anyVar: any;
anyVar();
anyVar.method();
```

```ts
// Right
declare const fn: () => void;
fn();
```

---

## 29. no-unsafe-enum-comparison **fix**

Disallow comparing an enum value with a non-enum value.

```ts
// Wrong
enum Fruit { Apple }
declare let fruit: Fruit;
fruit === 0;
```

```ts
// Right (fixed: use enum member)
enum Fruit { Apple }
declare let fruit: Fruit;
fruit === Fruit.Apple;
```

---

## 30. no-unsafe-member-access

Disallow member access on a value with type `any`.

```ts
// Wrong
declare const anyVar: any;
anyVar.foo;
anyVar['bar'];
```

```ts
// Right
declare const obj: { foo: string };
obj.foo;
```

---

## 31. no-unsafe-return

Disallow returning a value with type `any` from a function.

```ts
// Wrong
function foo(): string[] {
  return [] as any;
}
```

```ts
// Right
function foo(): string[] {
  return ['hello'];
}
```

---

## 32. no-unsafe-type-assertion

Disallow type assertions that narrow a type.

```ts
// Wrong
declare const x: string | number;
const y = x as number;
```

```ts
// Right
declare const x: string | number;
if (typeof x === 'number') {
  const y = x; // narrowed via type guard
}
```

---

## 33. no-unsafe-unary-minus

Require unary negation to take a number.

```ts
// Wrong
declare const s: string;
const x = -s;
```

```ts
// Right
declare const n: number;
const x = -n;
```

---

## 34. no-useless-default-assignment **fix**

Disallow default values that will never be used because the parameter is never undefined.

```ts
// Wrong
function foo({ bar = '' }: { bar: string }) {}
```

```ts
// Right (fixed: remove default since bar is required)
function foo({ bar }: { bar: string }) {}
```

---

## 35. non-nullable-type-assertion-style **fix**

Enforce non-null assertions over explicit type assertions.

```ts
// Wrong
declare const x: string | undefined;
const y = x as string;
```

```ts
// Right (fixed: use non-null assertion)
declare const x: string | undefined;
const y = x!;
```

---

## 36. only-throw-error

Disallow throwing non-Error values as exceptions.

```ts
// Wrong
throw 'error';
throw { message: 'error' };
```

```ts
// Right
throw new Error('error');
```

---

## 37. prefer-find **fix**

Enforce `Array.prototype.find()` over `.filter()[0]`.

```ts
// Wrong
[1, 2, 3].filter(x => x > 1)[0];
```

```ts
// Right (fixed: use find)
[1, 2, 3].find(x => x > 1);
```

---

## 38. prefer-includes **fix**

Enforce `includes` method over `indexOf` method.

```ts
// Wrong
declare const arr: string[];
arr.indexOf('foo') !== -1;
```

```ts
// Right (fixed: use includes)
declare const arr: string[];
arr.includes('foo');
```

---

## 39. prefer-nullish-coalescing **fix**

Enforce using the nullish coalescing operator instead of logical OR for nullable values.

```ts
// Wrong
declare const a: string | null;
const b = a || 'default';
```

```ts
// Right (fixed: use ??)
declare const a: string | null;
const b = a ?? 'default';
```

---

## 40. prefer-optional-chain **fix**

Enforce using concise optional chain expressions instead of chained logical ands.

```ts
// Wrong
foo && foo.a && foo.a.b;
```

```ts
// Right (fixed: use optional chain)
foo?.a?.b;
```

---

## 41. prefer-promise-reject-errors

Require using Error objects as Promise rejection reasons.

```ts
// Wrong
Promise.reject('error');
```

```ts
// Right
Promise.reject(new Error('error'));
```

---

## 42. prefer-readonly **fix**

Require private members to be marked as `readonly` if never modified outside the constructor.

```ts
// Wrong
class Foo {
  private name = 'bar';
}
```

```ts
// Right (fixed: add readonly)
class Foo {
  private readonly name = 'bar';
}
```

---

## 43. prefer-readonly-parameter-types

Require function parameters to be typed as `readonly` to prevent accidental mutation of inputs.

```ts
// Wrong
function foo(arr: string[]) {
  console.log(arr);
}
```

```ts
// Right
function foo(arr: readonly string[]) {
  console.log(arr);
}
```

---

## 44. prefer-reduce-type-parameter **fix**

Enforce using type parameter when calling `Array#reduce` instead of a type assertion.

```ts
// Wrong
[1, 2].reduce((acc, n) => [...acc, n], [] as number[]);
```

```ts
// Right (fixed: use type parameter)
[1, 2].reduce<number[]>((acc, n) => [...acc, n], []);
```

---

## 45. prefer-regexp-exec **fix**

Enforce `RegExp#exec` over `String#match` if no global flag is provided.

```ts
// Wrong
'something'.match(/thing/);
```

```ts
// Right (fixed: use exec)
/thing/.exec('something');
```

---

## 46. prefer-return-this-type **fix**

Enforce that `this` is used when only `this` type is returned.

```ts
// Wrong
class Foo {
  bar(): Foo { return this; }
}
```

```ts
// Right (fixed: return this type)
class Foo {
  bar(): this { return this; }
}
```

---

## 47. prefer-string-starts-ends-with **fix**

Enforce using `String#startsWith` and `String#endsWith` over other equivalent methods.

```ts
// Wrong
declare const s: string;
s.indexOf('bar') === 0;
/^bar/.test(s);
```

```ts
// Right (fixed: use startsWith)
declare const s: string;
s.startsWith('bar');
```

---

## 48. promise-function-async **fix**

Require any function or method that returns a Promise to be marked async.

```ts
// Wrong
function getData() {
  return Promise.resolve('value');
}
```

```ts
// Right (fixed: add async)
async function getData() {
  return Promise.resolve('value');
}
```

---

## 49. related-getter-setter-pairs

Enforce that `get()` types should be assignable to their equivalent `set()` type.

```ts
// Wrong
interface Box {
  get value(): string;
  set value(newValue: number);
}
```

```ts
// Right
interface Box {
  get value(): string;
  set value(newValue: string);
}
```

---

## 50. require-array-sort-compare

Require `Array#sort` and `Array#toSorted` calls to always provide a `compareFunction`.

```ts
// Wrong
declare const nums: number[];
nums.sort();
```

```ts
// Right
declare const nums: number[];
nums.sort((a, b) => a - b);
```

---

## 51. require-await **fix**

Disallow async functions which do not return promises and have no `await` expression.

```ts
// Wrong
async function returnNumber() {
  return 1;
}
```

```ts
// Right (fixed: remove async)
function returnNumber() {
  return 1;
}
```

---

## 52. restrict-plus-operands

Require both operands of addition to be the same type and be `bigint`, `number`, or `string`.

```ts
// Wrong
const a = 1n + 1;
const b = 'foo' + 42;
```

```ts
// Right
const a = 1n + 1n;
const b = 'foo' + String(42);
```

---

## 53. restrict-template-expressions

Enforce template literal expressions to be of `string` type.

```ts
// Wrong
const obj = { name: 'Foo' };
const msg = `arg = ${obj}`;
```

```ts
// Right
const obj = { name: 'Foo' };
const msg = `arg = ${obj.name}`;
```

---

## 54. return-await **fix**

Enforce consistent awaiting of returned promises (especially in try/catch).

```ts
// Wrong
async function foo() {
  try { return Promise.resolve('try'); }
  catch (e) { /* never executes! */ }
}
```

```ts
// Right (fixed: add await in try)
async function foo() {
  try { return await Promise.resolve('try'); }
  catch (e) { /* catches as expected */ }
}
```

---

## 55. strict-boolean-expressions **fix**

Disallow certain types in boolean expressions.

```ts
// Wrong
declare const num: number | undefined;
if (num) { console.log('truthy'); }
```

```ts
// Right (fixed: explicit check)
declare const num: number | undefined;
if (num != null) { console.log('defined'); }
```

---

## 56. strict-void-return

Disallow passing a value-returning function in a position accepting a void function.

```ts
// Wrong
declare function run(cb: () => void): void;
run(() => 42);
```

```ts
// Right
declare function run(cb: () => void): void;
run(() => { console.log('done'); });
```

---

## 57. switch-exhaustiveness-check **fix**

Require switch-case statements to be exhaustive.

```ts
// Wrong
type Day = 'Mon' | 'Tue' | 'Wed';
declare const day: Day;
switch (day) {
  case 'Mon': break;
  // missing Tue, Wed!
}
```

```ts
// Right (fixed: add missing cases)
type Day = 'Mon' | 'Tue' | 'Wed';
declare const day: Day;
switch (day) {
  case 'Mon': break;
  case 'Tue': break;
  case 'Wed': break;
}
```

---

## 58. unbound-method

Enforce unbound methods are called with their expected scope.

```ts
// Wrong
class Foo {
  log() { console.log(this); }
}
const instance = new Foo();
const myLog = instance.log; // unbound!
```

```ts
// Right
class Foo {
  log = () => { console.log(this); };
}
const instance = new Foo();
const myLog = instance.log; // bound via arrow
```

---

## 59. use-unknown-in-catch-callback-variable **fix**

Enforce typing arguments in Promise rejection callbacks as `unknown`.

```ts
// Wrong
promise.catch(err => {
  console.log(err.message);
});
```

```ts
// Right (fixed: type as unknown)
promise.catch((err: unknown) => {
  if (err instanceof Error) console.log(err.message);
});
```
