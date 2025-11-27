# Nexus

Nexus is a real-time collaborative task management application inspired by Trello.
My goal with this project is to showcase my ability to build and work with things at scale.
I have written a comprehensive feature list based on Trello's functionality and my goal is to checkmark as many as I can.

***

# **Nexus - Trello Feature Specification**

---

## **🎯 Core Entities & Data Models**

### **1. User**
- ID, email, password (hashed), name, username, avatar URL
- Email verification status
- OAuth account connections (Google, GitHub, etc.)
- Created/updated timestamps

### **2. Board**
- ID, name, description, owner ID
- Visibility: `private`, `team`, `public`
- Background (color/image)
- Members (user IDs with roles: `admin`, `normal`, `observer`)
- Settings (permissions, notifications)
- Archived status
- Created/updated timestamps

### **3. List**
- ID, name, board ID
- Position (order within board)
- Color (optional)
- Collapsed status
- Archived status
- Created/updated timestamps

### **4. Card**
- ID, title, description, list ID
- Position (order within list)
- Due date and time
- Start date (optional)
- Cover image/color
- Archived status
- Completed status
- Created/updated timestamps

### **5. Card Member**
- Card ID, User ID (many-to-many relationship)
- Assigned date

### **6. Label**
- ID, name, color
- Board ID (labels are board-specific)

### **7. Card Label**
- Card ID, Label ID (many-to-many relationship)

### **8. Checklist**
- ID, name, card ID
- Position (order within card)
- Created/updated timestamps

### **9. Checklist Item**
- ID, checklist ID, text
- Completed status
- Position (order within checklist)
- Due date (optional)
- Assigned to (user ID, optional)
- Created/updated timestamps

### **10. Attachment**
- ID, card ID, file name, file URL/path
- File type, file size
- Uploaded by (user ID)
- Created timestamp

### **11. Comment**
- ID, card ID, user ID, content
- Mentions (@username)
- Edited status and timestamp
- Created/updated timestamps

### **12. Activity/Event**
- ID, type (card_created, card_moved, comment_added, etc.)
- Board ID, card ID (optional), list ID (optional)
- User ID (who performed the action)
- Metadata (what changed, old/new values)
- Timestamp

### **13. Board Member**
- Board ID, User ID
- Role: `admin`, `normal`, `observer`
- Invited at, joined at timestamps

### **14. Card Vote**
- Card ID, User ID (many-to-many relationship)
- Voted at timestamp

### **15. Sticker**
- ID, card ID, name, image URL
- Position (x, y coordinates on card)
- Created timestamp

### **16. Power-Up**
- ID, name, description, enabled status
- Board ID (power-ups are board-specific)

### **17. Automation/Butler Rule**
- ID, board ID, name
- Trigger (event that starts the automation)
- Action (what happens when triggered)
- Enabled status
- Created/updated timestamps

### **18. Board Template**
- ID, name, description
- Template structure (JSON)
- Created by (user ID)
- Public/private status
- Created timestamp

### **19. View**
- ID, board ID, type (`calendar`, `timeline`, `table`, `dashboard`, `map`)
- Settings (JSON for view-specific configuration)
- Created/updated timestamps

### **20. Notification**
- ID, user ID, type
- Board ID, card ID (optional)
- Read status
- Created timestamp

---

## **📋 Phase 1: Foundation (MVP)**
*Goal: Basic CRUD + Authentication - Get it working*

### **Authentication & Authorization**

**User registration with email validation**
- [x] Credentials
- [x] OAuth

**Login with JWT token generation**
- [x] Credentials
- [x] OAuth

- [x] Token refresh mechanism
- [x] Password reset flow (email-based)
- [x] Logout (token invalidation)
- [x] Middleware: Protect routes, verify tokens
- [ ] Role-based access control (RBAC) for board members

### **User Management**
- [ ] Get user profile
- [ ] Update profile (name, username, avatar)
- [ ] Change password
- [ ] Delete account
- [ ] User preferences (theme, notifications, language)

### **Board Management**
- [x] Create board
- [ ] Get board details
- [x] Update board (name, description, background, visibility)
- [x] Delete board (owner only)
- [x] Archive/unarchive board
- [x] List user's boards (with filters: owned, member, archived)
- [x] Board visibility settings (private, team, public)
- [ ] Board background customization (colors and images)
- [ ] Copy/duplicate board
- [ ] Create board from template

### **List Management**
- [x] Create list within board
- [ ] Get list details
- [x] Update list (name, position, color)
- [x] Delete list
- [x] Archive/unarchive list
- [ ] Reorder lists (drag and drop position)
- [x] Collapse/expand list
- [ ] List color customization

### **Card Management (CRUD)**
- [x] Create card within list
- [x] Get card details (full card back view)
- [x] Update card (title, description, position)
- [x] Delete card
- [x] Archive/unarchive card
- [ ] Move card between lists
- [ ] Reorder cards within list (drag and drop)
- [ ] Copy/duplicate card
- [ ] Convert card to template

### **Card Details**
- [ ] Add/edit card description (rich text support)
- [ ] Add/edit due date and time
- [x] Add start date
- [ ] Mark due date as complete
- [x] Set card cover (color or image)
- [x] Remove card cover
- [ ] Card aging (visual indication of old cards)

### **Card Members**
- [ ] Add member to card
- [ ] Remove member from card
- [ ] List card members
- [ ] Multiple members per card support

### **Labels**
- [ ] Create label on board
- [ ] Update label (name, color)
- [ ] Delete label
- [ ] Add label to card
- [ ] Remove label from card
- [ ] List all labels on board
- [ ] List labels on specific card

### **Checklists**
- [ ] Create checklist on card
- [ ] Update checklist name
- [ ] Delete checklist
- [ ] Reorder checklists on card
- [ ] Add item to checklist
- [ ] Update checklist item (text, due date, assignee)
- [ ] Delete checklist item
- [ ] Toggle checklist item completion
- [ ] Reorder checklist items
- [ ] Convert checklist items to cards

### **Attachments**
- [ ] Upload attachment to card (images, PDFs, docs, etc.)
- [ ] Store files (local storage or S3/CloudFlare R2)
- [ ] List attachments on card
- [ ] Delete attachment
- [ ] Download attachment
- [ ] Preview attachments (images, PDFs)
- [ ] Attachment size limits and validation

### **Comments**
- [ ] Add comment to card
- [ ] Edit own comments
- [ ] Delete own comments
- [ ] List comments on card (with pagination)
- [ ] Mention users in comments (@username)
- [ ] Real-time comment updates

### **Basic API Features**
- [ ] Pagination for list endpoints
- [ ] Sorting (by position, created date, due date, name)
- [ ] Search cards by title/description across boards
- [ ] Error handling with proper HTTP status codes
- [ ] Request validation middleware
- [ ] Structured JSON responses

---

## **⚡ Phase 2: Real-Time Features**
*Goal: Add WebSocket support for live updates*

### **WebSocket Server**
- [ ] WebSocket connection handler
- [ ] Authentication via JWT in WebSocket handshake
- [ ] Connection management (store active connections per board)
- [ ] Heartbeat/ping-pong to detect disconnects
- [ ] Graceful connection cleanup

### **Real-Time Events Broadcasting**
- [ ] **Card Created** → Broadcast to board members
- [ ] **Card Updated** → Broadcast to board members
- [ ] **Card Moved** → Broadcast to board members
- [ ] **Card Archived** → Broadcast to board members
- [ ] **List Created** → Broadcast to board members
- [ ] **List Updated** → Broadcast to board members
- [ ] **List Moved** → Broadcast to board members
- [ ] **Comment Added** → Notify card watchers
- [ ] **Member Added to Card** → Notify assignee
- [ ] **Due Date Added/Changed** → Notify card members
- [ ] **Label Added/Removed** → Broadcast to board members
- [ ] **Checklist Item Completed** → Broadcast to board members
- [ ] **User Online/Offline** → Broadcast presence to board
- [ ] **User Viewing Card** → Show "👁️ User is viewing this card"

### **Presence System**
- [ ] Track which users are online in board
- [ ] Track which card a user is currently viewing
- [ ] Broadcast presence changes (online/offline/viewing)
- [ ] Show active users in board
- [ ] Show who's viewing each card

### **Notification System**
- [ ] In-app notification center (unread count)
- [ ] Mark notification as read
- [ ] Mark all notifications as read
- [ ] Clear all notifications
- [ ] Notification types:
  - Card assigned to you
  - Mentioned in comment
  - Card due date approaching
  - Card due date passed
  - Comment added to your card
  - Card moved to list you're watching
  - Added to board
  - Removed from board

---

## **🔥 Phase 3: Collaboration Features**
*Goal: Make it truly collaborative*

### **Board Members & Permissions**
- [x] Invite users to board (via email)
- [x] Accept/decline board invitations
- [ ] Remove member from board
- [ ] Change member role (admin, normal, observer)
- [ ] List board members with online status
- [ ] Board member permissions:
  - Admin: Full control
  - Normal: Can edit
  - Observer: Read-only

### **Card Watching**
- [ ] Watch card (get notifications for all updates)
- [ ] Unwatch card
- [ ] List cards user is watching
- [ ] Auto-watch when assigned to card

### **Card Voting**
- [ ] Vote on card
- [ ] Remove vote from card
- [ ] List votes on card (who voted)
- [ ] Vote count display

### **Stickers**
- [ ] Add sticker to card
- [ ] Remove sticker from card
- [ ] Move sticker on card (position)
- [ ] List stickers on card
- [ ] Sticker library/presets

### **Activity Feed**
- [ ] Board activity feed (all recent actions)
- [ ] Card activity timeline (who did what, when)
- [ ] Filter activity by action type
- [ ] Filter activity by user
- [ ] Real-time activity updates

### **Board Templates**
- [ ] Create board template
- [ ] Use template to create board
- [ ] List public templates
- [ ] List user's templates
- [ ] Share template
- [ ] Delete template

---

## **🚀 Phase 4: Advanced Features**
*Goal: Production-ready, enterprise-level functionality*

### **Views**
- [ ] **Calendar View**: View cards by due date
- [ ] **Timeline View**: Gantt-style timeline of cards
- [ ] **Table View**: Spreadsheet-like view of cards
- [ ] **Dashboard View**: Analytics and metrics
- [ ] **Map View**: Location-based cards (if location data added)
- [ ] Switch between views
- [ ] Save view preferences per board

### **Search & Filters**
- [ ] Full-text search across cards, comments, descriptions
- [ ] Advanced filters:
  - Filter by labels (multiple, AND/OR logic)
  - Filter by members (multiple)
  - Filter by due date (overdue, due today, due this week, custom range)
  - Filter by completed status
  - Filter by archived status
  - Filter by board
  - Filter by list
- [ ] Save filter presets
- [ ] Search suggestions/autocomplete
- [ ] Search within specific board or all boards

### **Automation (Butler Rules)**
- [ ] Create automation rule
- [ ] Edit automation rule
- [ ] Delete automation rule
- [ ] Enable/disable automation
- [ ] Automation triggers:
  - Card created
  - Card moved to list
  - Card due date approaching
  - Card completed
  - Checklist item completed
  - Member added to card
  - Comment added
- [ ] Automation actions:
  - Move card to list
  - Add label to card
  - Remove label from card
  - Assign member to card
  - Remove member from card
  - Set due date
  - Add checklist
  - Archive card
  - Create card
- [ ] List automations on board

### **Power-Ups (Integrations)**
- [ ] Enable power-up on board
- [ ] Disable power-up on board
- [ ] List available power-ups
- [ ] Power-up settings/config
- [ ] Popular power-ups:
  - Calendar integration
  - Voting
  - Custom fields
  - Time tracking
  - Reporting/analytics

### **Performance & Scalability**
- [ ] **Redis Caching**:
  - Cache frequently accessed data (user sessions, board details)
  - Cache card lists (invalidate on update)
  - Cache board members
- [ ] **Rate Limiting**:
  - Per-user rate limits (100 requests/min)
  - Per-IP rate limits for public endpoints
  - WebSocket message rate limiting
- [ ] **Database Optimization**:
  - Indexes on frequently queried fields
  - Query optimization (use EXPLAIN ANALYZE)
  - Connection pooling
  - Efficient position updates for drag-and-drop

### **Background Jobs**
- [ ] Email notifications (card assigned, mentions, due dates)
- [ ] Daily digest emails (cards due today, overdue)
- [ ] Weekly summary emails
- [ ] Cleanup expired tokens
- [ ] Archive old completed cards (configurable)
- [ ] Generate board reports

### **Analytics & Insights**
- [ ] Cards completed per board
- [ ] Average time to complete card
- [ ] Overdue cards report
- [ ] User activity metrics
- [ ] Board activity metrics
- [ ] List performance (cards moved through)
- [ ] Member contribution stats
- [ ] Export analytics as CSV/JSON

### **Advanced Card Features**
- [ ] Card templates
- [ ] Bulk operations (move multiple cards, assign multiple, add labels)
- [ ] Custom fields per board (text, number, date, dropdown, checkbox)
- [ ] Card linking (link to another card)
- [ ] Card merging
- [ ] Card aging (visual fade for old cards)

---

## **🛡️ Phase 5: Security & DevOps**
*Goal: Make it production-ready and secure*

### **Security**
- [ ] **Input Validation**: Sanitize all user inputs
- [ ] **SQL Injection Prevention**: Use parameterized queries
- [ ] **XSS Prevention**: Escape outputs if serving HTML
- [ ] **CSRF Protection**: For any state-changing operations
- [ ] **Rate Limiting**: Prevent brute force/DoS
- [ ] **Secure Headers**: CORS, CSP, HSTS
- [ ] **Audit Logging**: Log sensitive actions (login, role changes, board deletion)
- [ ] **File Upload Security**: Validate file types, scan for malware
- [ ] **Board Access Control**: Verify user has access to board before operations

### **Testing**
- [ ] Unit tests for business logic (80%+ coverage)
- [ ] Integration tests for API endpoints
- [ ] WebSocket connection tests
- [ ] Load testing (simulate 1000+ concurrent users)
- [ ] Database migration tests
- [ ] E2E tests for critical user flows

### **DevOps & Deployment**
- [ ] **Docker**: Containerize the application
- [ ] **Docker Compose**: Multi-container setup (app + postgres + redis)
- [ ] **Environment Config**: Use `.env` files, never hardcode secrets
- [ ] **CI/CD Pipeline** (GitHub Actions):
  - Run tests on every push
  - Lint code (golangci-lint)
  - Build Docker image
  - Deploy to staging/production
- [ ] **Monitoring**:
  - Prometheus metrics endpoint
  - Grafana dashboards (request rate, error rate, latency)
  - Logging (structured logs with zerolog or zap)
- [ ] **Database Migrations**: Use `golang-migrate` or similar
- [ ] **Health Check Endpoint**: `/health` for load balancers

### **Documentation**
- [ ] **API Documentation**: Swagger/OpenAPI spec
- [ ] **README**: Setup instructions, architecture overview
- [ ] **Architecture Diagram**: Show components (API, WebSocket, DB, Redis)
- [ ] **Deployment Guide**: How to deploy to production
- [ ] **Contribution Guide**: If open-sourcing

---

## **🎨 Phase 6: Polish & Extras (Optional)**
*Goal: Stand out from other candidates*

### **Nice-to-Have Features**
- [ ] **Dark Mode API Support**: Send theme preferences
- [ ] **Card Time Tracking**: Start/stop timer on cards
- [ ] **Card Linking**: Link cards together
- [ ] **Webhooks**: Trigger external services on events
- [ ] **API Versioning**: `/api/v1/` for future-proofing
- [ ] **GraphQL API**: Alternative to REST (bonus points)
- [ ] **Mobile Push Notifications**: FCM/APNS integration
- [ ] **Board Export**: Export board as CSV/PDF/JSON
- [ ] **Slack Integration**: Post updates to Slack channels
- [ ] **Email to Card**: Create card from email
- [ ] **Card Location**: Add location/address to cards
- [ ] **Card Aging**: Visual fade for cards not updated in X days

### **Developer Experience**
- [ ] **Postman Collection**: Pre-configured API requests
- [ ] **Sample Data Seeder**: Populate DB with test data
- [ ] **CLI Tool**: Manage boards/cards from terminal
- [ ] **SDK/Client Library**: Go client for the API

---

## **📊 Feature Priority Matrix**

| **Must Have (Phase 1-2)** | **Should Have (Phase 3)** | **Nice to Have (Phase 4-6)** |
|---------------------------|---------------------------|------------------------------|
| Authentication | Comments & Mentions | Analytics Dashboard |
| Board/List/Card CRUD | File Attachments | Recurring Cards |
| WebSocket Real-time | Activity Feed | GraphQL API |
| Presence System | Board Members | Webhooks |
| Basic Notifications | Labels & Checklists | Time Tracking |
| Drag & Drop (Position) | Card Voting | Power-Ups |
| Card Details | Board Templates | Multiple Views |

---

## **🎯 Recommended Build Order**

**Week 1-2:** Phase 1 (Foundation)  
**Week 3:** Phase 2 (WebSocket + Real-time)  
**Week 4:** Phase 3 (Collaboration)  
**Week 5:** Phase 4 (Advanced Features, Performance)  
**Week 6:** Phase 5 (Testing, DevOps, Documentation)  
**Week 7+:** Phase 6 (Extras if time permits)

---

## **Next Steps**

Now that we have the full Trello-inspired feature spec:

1. **Database schema updated** to match Trello structure (Board -> List -> Card)
2. **API endpoints** need to be designed for all features
3. **WebSocket events** need to be defined for real-time updates
4. **Frontend components** need to be built for drag-and-drop interface

Let me know which part you want to dive into next! 🚀
